// Package mttacc provides a model for retrieving and manipulating Mintter Account Objects.
package mttacc

import (
	"context"
	"fmt"
	"mintter/backend/core"
	"mintter/backend/vcs/vcsdb"
	"mintter/backend/vcs/vcstypes"
	"time"

	"github.com/ipfs/go-cid"
)

// Account-related attributes.
const (
	AttrAlias        vcsdb.Attribute = "mintter.account/alias"
	AttrEmail        vcsdb.Attribute = "mintter.account/email"
	AttrBio          vcsdb.Attribute = "mintter.account/bio"
	AttrRegistration vcsdb.Attribute = "mintter.account/registration"
	AttrProof        vcsdb.Attribute = "mintter.account.registration/proof"
	AttrDevice       vcsdb.Attribute = "mintter.account.registration/device"
)

// Register links device under a given account. Returns the CID of the resulting account object.
func Register(ctx context.Context, acc, device core.KeyPair, conn *vcsdb.Conn) (c cid.Cid, err error) {
	aid := acc.CID()

	perma, err := vcsdb.NewPermanode(vcstypes.NewAccountPermanode(aid))
	if err != nil {
		return c, err
	}

	obj := conn.NewObject(perma)
	id := conn.EnsureIdentity(core.NewIdentity(acc.PublicKey, device))
	change := conn.NewChange(obj, id, nil, time.Now().UTC())
	newDatom := vcsdb.MakeDatomFactory(change, 1, 0)

	proof, err := vcstypes.NewRegistrationProof(acc, device.CID())
	if err != nil {
		return c, err
	}

	reg := vcsdb.NewNodeID()
	conn.AddDatom(obj, newDatom(reg, AttrDevice, device.CID()))
	conn.AddDatom(obj, newDatom(reg, AttrProof, []byte(proof)))
	conn.AddDatom(obj, newDatom(vcsdb.RootNode, AttrRegistration, reg))

	conn.SaveVersion(obj, "main", id, vcsdb.LocalVersion{change})
	conn.EncodeChange(change, device)

	return perma.ID, nil
}

// GetDeviceProof searches for a registration proof of a device under an account.
func GetDeviceProof(conn *vcsdb.Conn, me core.Identity, account, device cid.Cid) (proof []byte, err error) {
	perma, err := vcsdb.NewPermanode(vcstypes.NewAccountPermanode(account))
	if err != nil {
		return nil, err
	}

	obj := conn.LookupPermanode(perma.ID)
	localMe := conn.EnsureIdentity(me)
	ver := conn.GetVersion(obj, "main", localMe)
	cs := conn.ResolveChangeSet(obj, ver)

	regs := conn.QueryValuesByAttr(obj, cs, vcsdb.RootNode, AttrRegistration)
	for _, reg := range regs {
		dd := conn.QueryLastValue(obj, cs, reg.Value.(vcsdb.NodeID), AttrDevice)
		if !dd.Value.(cid.Cid).Equals(device) {
			continue
		}
		proof := conn.QueryLastValue(obj, cs, reg.Value.(vcsdb.NodeID), AttrProof)
		return proof.Value.([]byte), nil
	}

	return nil, fmt.Errorf("proof not found")
}