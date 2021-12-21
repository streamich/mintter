package backend

import (
	"context"
	"encoding/hex"
	"fmt"
	p2p "mintter/backend/api/p2p/v1alpha"
	"mintter/backend/lndhub"
	"mintter/backend/wallet"
	"strings"

	"github.com/ipfs/go-cid"
	"go.uber.org/zap"
)

type lnclient struct {
	Lndhub *lndhub.Client
	Lnd    interface{} // TODO: implement LND client
}
type InvoiceRequest struct {
	AmountSats   int64  `help:"The invoice amount in satoshis" default:"0"`
	Memo         string `help:"Optional requested memo to be attached in the invoice" default:""`
	HoldInvoice  bool   `help:"If we request a hold invoice instead of a regular one. If true, then the following field is mandatory" default:"false"`
	PreimageHash []byte `help:"Preimage hash of the requested hold invoice. If HoldInvoice flag is set to false this field is skipped" default:""`
}

// RemoteInvoiceRequest requests a remote account to issue an invoice so we can pay it.
// Any of the devices associated with the remote account can issue it. For each
// associated device we found online ,we ask if it can provide an invoice.
// If for some reason, that device cannot create the invoice (insufficient
// inbound liquidity) we ask the next device. We return in the first device that
// can issue the invoice. If none of them can, then an error is raised.
func (srv *backend) RemoteInvoiceRequest(ctx context.Context, account AccountID,
	request InvoiceRequest) (string, error) {
	if _, err := srv.readyIPFS(); err != nil {
		srv.log.Warn("Account not ready yet", zap.String("Mintter account ID", account.String()), zap.String("Error", err.Error()))
		return "", fmt.Errorf("account %s not ready yet.", account.String())
	}

	if account.Equals(srv.repo.acc.id) {
		return "", fmt.Errorf("cannot remotely issue an invoice to myself")
	}

	all, err := srv.db.ListAccountDevices(ctx)
	if err != nil {
		srv.log.Warn("couldn't list devices", zap.String("Mintter account ID", account.String()), zap.String("Error", err.Error()))
		return "", fmt.Errorf("couldn't list devices from account ID %s.", account.String())
	}

	if devices, found := all[account]; found {
		for _, deviceID := range devices {
			cc, err := srv.dialPeer(ctx, deviceID.PeerID())
			if err != nil {
				continue
			}
			p2pc := p2p.NewP2PClient(cc)

			remoteInvoice, err := p2pc.RequestInvoice(ctx, &p2p.RequestInvoiceRequest{
				AmountSats:   request.AmountSats,
				Memo:         request.Memo,
				HoldInvoice:  request.HoldInvoice,
				PreimageHash: request.PreimageHash,
			})

			if err != nil {
				srv.log.Warn("coulnd't request invoice", zap.String("Error", err.Error()))
				return "", fmt.Errorf("request invoice failed.")
			}

			if remoteInvoice.PayReq == "" {
				return "", fmt.Errorf("received an empty invoice from remote peer")
			}

			return remoteInvoice.PayReq, nil
		}
		return "", fmt.Errorf("none of the devices associated with the provided account were reachable")
	}

	return "", fmt.Errorf("couln't find account %s", account.String())
}

// InsertWallet first tries to connect to the wallet with the provided credentials. On
// success, gets the wallet balance and inserts all that information in the database.
// InsertWallet returns the wallet actually inserted on success. The credentias are stored
// in plain text at the moment.
func (srv *backend) InsertWallet(ctx context.Context, walletType, credentialsURL, name string) (wallet.Wallet, error) {
	var err error
	var ret wallet.Wallet
	if strings.ToLower(walletType) != lndhub.LndhubWalletType { // TODO: support LND wallets
		return ret, fmt.Errorf(" wallet typy not supported. currently only %s", lndhub.LndhubWalletType)
	}

	creds, err := lndhub.ParseCredentials(credentialsURL)
	if err != nil {
		return ret, err
	}

	// Trying to authenticate with the provided credentials
	creds.Token, err = srv.lightningClient.Lndhub.Auth(ctx, creds)
	if err != nil {
		srv.log.Warn("couldn't authenticate new wallet", zap.String("Error", err.Error()))
		return ret, fmt.Errorf("couldn't authenticate new wallet %s. Please check provided credentials", name)
	}

	balanceSats, err := srv.lightningClient.Lndhub.GetBalance(ctx, creds)
	if err != nil {
		return ret, err
	}
	conn := srv.pool.Get(ctx)
	if conn == nil {
		return ret, fmt.Errorf("couldn't get sqlite connector from the pool before timeout. New wallet %s has not been inserted in database", name)
	}
	defer srv.pool.Put(conn)

	ret.Address = creds.ConnectionURL
	ret.Balance = int64(balanceSats)
	ret.ID = creds.ID
	ret.Name = name
	ret.Type = lndhub.LndhubWalletType
	binaryToken, err := hex.DecodeString(creds.Token) // TODO: encrypt the token before storing

	if err != nil {
		srv.log.Warn("could not decode Token", zap.String("Error", err.Error()))
		return ret, fmt.Errorf("couldn't decode token before insert the wallet in the database.")
	}

	if err = wallet.InsertWallet(conn, ret, binaryToken); err != nil {
		srv.log.Warn("couldn't insert wallet in the database", zap.String("Error", err.Error()))
		if strings.Contains(err.Error(), wallet.AlreadyExistsError) {
			return ret, fmt.Errorf("couldn't insert wallet %s in the database. ID already exists", name)
		}
		return ret, fmt.Errorf("couldn't insert wallet %s in the database. Unknown reason", name)
	}

	return ret, err
}

// ListWallets returns all the wallets available in the database.
func (srv *backend) ListWallets(ctx context.Context) ([]wallet.Wallet, error) {
	conn := srv.pool.Get(ctx)
	if conn == nil {
		return nil, fmt.Errorf("couldn't get sqlite connector from the pool before timeout")
	}
	defer srv.pool.Put(conn)
	wallets, err := wallet.ListWallets(conn, -1)
	if err != nil {
		srv.log.Warn("couldn't list wallets", zap.String("Error", err.Error()))
		return nil, fmt.Errorf("couldn't list wallets")
	}
	for _, w := range wallets {
		if strings.ToLower(w.Type) == lndhub.LndhubWalletType {
			token, err := wallet.GetAuth(conn, w.ID)
			if err != nil {
				srv.log.Warn("couldn't get auth", zap.String("Wallet", w.Name), zap.String("Error", err.Error()))
				return nil, fmt.Errorf("couldn't get auth from wallet %s.", w.Name)
			}
			creds := lndhub.Credentials{
				ConnectionURL: w.Address,
				Token:         hex.EncodeToString(token),
				ID:            w.ID,
			}
			balance, err := srv.lightningClient.Lndhub.GetBalance(ctx, creds)
			if err != nil {
				srv.log.Warn("couldn't get balance", zap.String("Wallet", w.Name), zap.String("Error", err.Error()))
				return nil, fmt.Errorf("couldn't get balance from wallet %s", w.Name)
			}
			w.Balance = int64(balance)
		}
	}
	return wallets, nil
}

// DeleteWallet removes the wallet given a valid ID string representing
// the url hash in case of Lndhub-type wallet or the pubkey in case of LND.
// If the removed wallet was the default wallet, a random wallet will be
// chosen as new default. Although it is advised that the user manually
// changes the default wallet after removing the previous default
func (srv *backend) DeleteWallet(ctx context.Context, walletID string) error {
	conn := srv.pool.Get(ctx)
	if conn == nil {
		return fmt.Errorf("couldn't get sqlite connector from the pool before timeout")
	}
	defer srv.pool.Put(conn)
	if err := wallet.RemoveWallet(conn, walletID); err != nil {
		srv.log.Warn("couldn't remove wallet", zap.String("WalletID", walletID), zap.String("Error", err.Error()))
		return fmt.Errorf("couldn't remove wallet %s", walletID)
	}
	return nil
}

// UpdateWalletName updates an existing wallet's name with the one provided.
// If the wallet represented by the id id does not exist, this function
// returns error. nil otherwise, along with the updated wallet.
func (srv *backend) UpdateWalletName(ctx context.Context, walletID string, newName string) (wallet.Wallet, error) {
	var ret wallet.Wallet
	var err error
	conn := srv.pool.Get(ctx)
	if conn == nil {
		return ret, fmt.Errorf("couldn't get sqlite connector from the pool before timeout")
	}
	defer srv.pool.Put(conn)
	if ret, err = wallet.UpdateWalletName(conn, walletID, newName); err != nil {
		srv.log.Warn("couldn't update wallet", zap.String("WalletID", walletID), zap.String("Error", err.Error()))
		return ret, fmt.Errorf("couldn't update wallet %s", walletID)
	}

	return ret, nil
}

// SetDefaultWallet sets the default wallet to the one that matches walletID.
// Previous default wallet is replaced by the new one so only one can be
// the default at any given time. The default wallet is the first wallet ever
// created until manually changed.
func (srv *backend) SetDefaultWallet(ctx context.Context, walletID string) (wallet.Wallet, error) {
	conn := srv.pool.Get(ctx)
	if conn == nil {
		return wallet.Wallet{}, fmt.Errorf("couldn't get sqlite connector from the pool before timeout")
	}
	defer srv.pool.Put(conn)

	defaultWallet, err := wallet.UpdateDefaultWallet(conn, walletID)
	if err != nil {
		srv.log.Warn("failed to update default wallet", zap.String("Error", err.Error()))
		return wallet.Wallet{}, fmt.Errorf("failed to update default wallet")
	}

	return defaultWallet, nil
}

// GetDefaultWallet gets the user's default wallet. If the user didn't manually
// update the default wallet, then the first wallet ever created is the default
// wallet. It will remain default until manually changed.
func (srv *backend) GetDefaultWallet(ctx context.Context) (wallet.Wallet, error) {
	conn := srv.pool.Get(ctx)
	if conn == nil {
		return wallet.Wallet{}, fmt.Errorf("couldn't get sqlite connector from the pool before timeout")
	}
	defer srv.pool.Put(conn)

	defaultWallet, err := wallet.GetDefaultWallet(conn)
	if err != nil {
		srv.log.Warn("failed to get default wallet", zap.String("Error", err.Error()))
		return wallet.Wallet{}, fmt.Errorf("failed to get default wallet")
	}

	return defaultWallet, nil

}

// RequestInvoice asks a remote peer to issue an invoice. Any of the devices associated with the accountID
// can issue the invoice. The memo field is optional and can be left nil
func (srv *backend) RequestInvoice(ctx context.Context, accountID string, amountSats int64, memo *string) (string, error) {
	conn := srv.pool.Get(ctx)
	if conn == nil {
		return "", fmt.Errorf("couldn't get sqlite connector from the pool before timeout")
	}
	defer srv.pool.Put(conn)

	invoiceMemo := ""
	if memo != nil {
		invoiceMemo = *memo
	}
	cID, err := cid.Decode(accountID)
	if err != nil {
		srv.log.Warn("couldn't parse accountID string", zap.String("Mintter Account ID", accountID), zap.String("Error", err.Error()))
		return "", fmt.Errorf("couldn't parse accountID string [%s], please check it is a proper accountID.", accountID)
	}

	payReq, err := srv.RemoteInvoiceRequest(ctx, AccountID(cID),
		InvoiceRequest{
			AmountSats:   amountSats,
			Memo:         invoiceMemo,
			HoldInvoice:  false,    // TODO: Add support hold invoices
			PreimageHash: []byte{}, // Only aplicable to hold invoices
		})
	if err != nil {
		return "", err
	}
	return payReq, nil
}

// PayInvoice tries to pay the provided invoice. If a walletID is provided, that wallet will be used instead of the default one
// If amountSats is provided, the invoice will be paid with that amount. This amount should be equal to the amount on the invoice
// unless the amount on the invoice is 0.
func (srv *backend) PayInvoice(ctx context.Context, payReq string, walletID *string, amountSats *uint64) (string, error) {
	var walletToPay wallet.Wallet
	var err error
	var amountToPay uint64
	conn := srv.pool.Get(ctx)

	if conn == nil {
		return "", fmt.Errorf("couldn't get sqlite connector from the pool before timeout")
	}
	defer srv.pool.Put(conn)

	if walletID != nil {
		walletToPay, err = wallet.GetWallet(conn, *walletID)
		if err != nil {
			srv.log.Warn("couldn't get wallet", zap.String("Wallet ID", *walletID), zap.String("Error", err.Error()))
			return "", fmt.Errorf("couldn't get wallet %s", *walletID)
		}
	} else {
		walletToPay, err = srv.GetDefaultWallet(ctx)
		if err != nil {
			srv.log.Warn("couldn't get default wallet", zap.String("Error", err.Error()))
			return "", fmt.Errorf("couldn't get default wallet")
		}
	}

	if walletToPay.Type != lndhub.LndhubWalletType {
		return "", fmt.Errorf("wallet type [%s] not supported to pay (yet)", walletToPay.Type)
	}

	if amountSats == nil {
		invoice, err := lndhub.DecodeInvoice(payReq)
		if err != nil {
			srv.log.Warn("couldn't decode invoice", zap.String("Error", err.Error()))
			return "", fmt.Errorf("couldn't decode invoice [%s], please make sure it is a bolt-11 complatible invoice", payReq)
		}
		amountToPay = uint64(invoice.MilliSat.ToSatoshis())
	} else {
		amountToPay = *amountSats
	}

	binaryToken, err := wallet.GetAuth(conn, walletToPay.ID)
	if err != nil {
		srv.log.Warn("couldn't get autorization to pay", zap.String("Wallet ID", walletToPay.ID), zap.String("Error", err.Error()))
		return "", fmt.Errorf("couldn't get autorization to pay with wallet ID [%s]", walletToPay.ID)
	}

	if err = srv.lightningClient.Lndhub.PayInvoice(ctx, lndhub.Credentials{
		ConnectionURL: walletToPay.Address,
		Token:         hex.EncodeToString(binaryToken),
	}, payReq, amountToPay); err != nil {
		srv.log.Warn("couldn't pay invoice", zap.String("Error", err.Error()))
		if strings.Contains(err.Error(), wallet.NotEnoughBalance) {
			return "", fmt.Errorf("couldn't pay invoice. Insufficient balance in wallet name %s", walletToPay.Name)
		}
		if strings.Contains(err.Error(), wallet.InvoiceQttyMissmatch) {
			return "", fmt.Errorf("couldn't pay invoice. %s", err.Error())
		}
		return "", fmt.Errorf("couldn't pay invoice. Unknown reason")
	}

	return walletToPay.ID, nil

}