package documents

import (
	"context"
	"fmt"
	"mintter/backend/core"
	"mintter/backend/crdt"
	documents "mintter/backend/genproto/documents/v1alpha"
	"mintter/backend/ipfs"
	"mintter/backend/vcs"
	"mintter/backend/vcs/vcssql"
	"mintter/backend/vcs/vcstypes"

	"crawshaw.io/sqlite/sqlitex"
	"github.com/ipfs/go-cid"
	cbornode "github.com/ipfs/go-ipld-cbor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type (
	DraftsServer       = documents.DraftsServer
	PublicationsServer = documents.PublicationsServer
	ContentGraphServer = documents.ContentGraphServer
)

type Server struct {
	me  core.Identity
	db  *sqlitex.Pool
	vcs *vcs.SQLite
}

func NewServer(me core.Identity, db *sqlitex.Pool, vcs *vcs.SQLite) *Server {
	return &Server{
		me:  me,
		db:  db,
		vcs: vcs,
	}
}

func (api *Server) CreateDraft(ctx context.Context, in *documents.CreateDraftRequest) (*documents.Document, error) {
	if in.ExistingDocumentId != "" {
		// Load time dag.
		// Create working copy.
		return nil, status.Errorf(codes.Unimplemented, "updating publications is not implemented yet")
	}

	me := api.me.AccountID()

	p := vcstypes.NewDocumentPermanode(me)

	permablk, err := vcs.EncodeBlock[vcs.Permanode](p)
	if err != nil {
		return nil, err
	}

	if err := api.vcs.StorePermanode(ctx, permablk.Block, permablk.Value); err != nil {
		return nil, err
	}

	wc := vcs.NewWorkingCopy(permablk.Cid(), "main")

	if err := api.vcs.SaveWorkingCopy(ctx, wc); err != nil {
		return nil, err
	}

	{
		conn, release, err := api.db.Conn(ctx)
		if err != nil {
			return nil, err
		}
		defer release()

		ocodec, ohash := ipfs.DecodeCID(permablk.Cid())

		if err := vcssql.DraftsInsert(conn, ohash, int(ocodec), "", "", int(p.CreateTime.Unix()), int(p.CreateTime.Unix())); err != nil {
			return nil, err
		}
	}

	return &documents.Document{
		Id:         permablk.Cid().String(),
		Author:     me.String(),
		CreateTime: timestamppb.New(p.CreateTime),
		UpdateTime: timestamppb.New(p.CreateTime),
	}, nil
}

func (api *Server) UpdateDraft(ctx context.Context, in *documents.UpdateDraftRequest) (*documents.Document, error) {
	return nil, status.Error(codes.Unimplemented, "deprecated")
}

func (api *Server) UpdateDraftV2(ctx context.Context, in *documents.UpdateDraftRequestV2) (*emptypb.Empty, error) {
	oid, err := cid.Decode(in.DocumentId)
	if err != nil {
		return nil, fmt.Errorf("failed to decode document id: %w", err)
	}

	if in.Changes == nil {
		return nil, status.Errorf(codes.InvalidArgument, "must send some changes to apply to the document")
	}

	draft, err := api.getDraft(ctx, oid, "main")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve old draft: %w", err)
	}

	doc := draft.doc
	wc := draft.wc

	for _, c := range in.Changes {
		switch op := c.Op.(type) {
		case *documents.DocumentChange_SetTitle:
			doc.ChangeTitle(op.SetTitle)
		case *documents.DocumentChange_SetSubtitle:
			doc.ChangeSubtitle(op.SetSubtitle)
		case *documents.DocumentChange_MoveBlock_:
			if err := doc.MoveBlock(op.MoveBlock.BlockId, op.MoveBlock.Parent, op.MoveBlock.LeftSibling); err != nil {
				return nil, err
			}
		case *documents.DocumentChange_ReplaceBlock:
			blk, err := blockFromProto(op.ReplaceBlock)
			if err != nil {
				return nil, err
			}
			if err := doc.ReplaceBlock(blk); err != nil {
				return nil, err
			}
		case *documents.DocumentChange_DeleteBlock:
			if err := doc.DeleteBlock(op.DeleteBlock); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("invalid draft update operation %T: %+v", c, c)
		}
	}

	oldEvents := draft.oldEvents
	newEvents := doc.Events()

	draftEvents := make([]vcstypes.DocumentEvent, len(oldEvents)+len(newEvents))
	n := copy(draftEvents, draft.oldEvents)
	copy(draftEvents[n:], doc.Events())

	data, err := cbornode.DumpObject(draftEvents)
	if err != nil {
		return nil, fmt.Errorf("failed to encode merged draft events: %w", err)
	}

	wc.SetData(data)

	if err := api.vcs.SaveWorkingCopy(ctx, wc); err != nil {
		return nil, fmt.Errorf("failed to save draft working copy: %w", err)
	}

	{
		conn, release, err := api.db.Conn(ctx)
		if err != nil {
			return nil, err
		}
		defer release()

		ocodec, ohash := ipfs.DecodeCID(oid)

		if err := vcssql.DraftsUpdate(conn, doc.State().Title, doc.State().Subtitle, int(doc.State().UpdateTime.Unix()), ohash, int(ocodec)); err != nil {
			return nil, err
		}
	}

	// TODO: index links.
	// Move old links insert new links.

	return &emptypb.Empty{}, nil
}

func (api *Server) GetDraft(ctx context.Context, in *documents.GetDraftRequest) (*documents.Document, error) {
	oid, err := cid.Decode(in.DocumentId)
	if err != nil {
		return nil, err
	}

	draft, err := api.getDraft(ctx, oid, "main")
	if err != nil {
		return nil, err
	}

	return docToProto(draft.doc)
}

func (api *Server) ListDrafts(ctx context.Context, in *documents.ListDraftsRequest) (*documents.ListDraftsResponse, error) {
	conn, release, err := api.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer release()

	res, err := vcssql.DraftsList(conn)
	if err != nil {
		return nil, err
	}

	out := &documents.ListDraftsResponse{
		Documents: make([]*documents.Document, len(res)),
	}

	aid := api.me.AccountID().String()

	for i, l := range res {
		out.Documents[i] = &documents.Document{
			Id:         cid.NewCidV1(uint64(l.ObjectsCodec), l.ObjectsMultihash).String(),
			Author:     aid,
			Title:      l.DraftsTitle,
			Subtitle:   l.DraftsSubtitle,
			CreateTime: &timestamppb.Timestamp{Seconds: int64(l.DraftsCreateTime)},
			UpdateTime: &timestamppb.Timestamp{Seconds: int64(l.DraftsUpdateTime)},
		}
	}

	return out, nil
}

func (api *Server) PublishDraft(ctx context.Context, in *documents.PublishDraftRequest) (*documents.Publication, error) {
	oid, err := cid.Decode(in.DocumentId)
	if err != nil {
		return nil, err
	}

	wc, err := api.vcs.LoadWorkingCopy(ctx, oid, "main")
	if err != nil {
		return nil, err
	}

	// TODO: ensure transactionality here.
	recorded, err := api.vcs.RecordChange(ctx, oid, api.me, wc.Version(), "mintter.Document", wc.Data())
	if err != nil {
		return nil, err
	}

	if err := api.vcs.RemoveWorkingCopy(ctx, oid, "main"); err != nil {
		return nil, err
	}

	newVer := vcs.NewVersion(recorded.LamportTime, recorded.ID)

	// TODO: implement optimistic concurrency control here.
	if err := api.vcs.StoreNamedVersion(ctx, oid, api.me, "main", newVer); err != nil {
		return nil, err
	}

	doc, err := api.getPublication(ctx, oid, newVer)
	if err != nil {
		return nil, err
	}

	docpb, err := docToProto(doc)
	if err != nil {
		return nil, err
	}
	docpb.PublishTime = timestamppb.New(recorded.CreateTime)

	pub := &documents.Publication{
		Version:  newVer.String(),
		Document: docpb,
		// TODO: get real latest version.
		LatestVersion: newVer.String(),
	}

	// TODO: move this elsewhere. Combine db writes into one transaction.
	{
		conn, release, err := api.db.Conn(ctx)
		if err != nil {
			return nil, err
		}
		defer release()

		ocodec, ohash := ipfs.DecodeCID(oid)

		if err := vcssql.DraftsDelete(conn, ohash, int(ocodec)); err != nil {
			return nil, err
		}

		if err := vcssql.PublicationsUpsert(conn, ohash, int(ocodec),
			doc.State().Title,
			doc.State().Subtitle,
			int(doc.State().CreateTime.Unix()),
			int(doc.State().UpdateTime.Unix()),
			int(pub.Document.PublishTime.Seconds),
			pub.LatestVersion,
		); err != nil {
			return nil, err
		}
	}

	// Delete draft from the index.
	// Add publication to the index.

	return pub, nil
}

func (api *Server) DeleteDraft(ctx context.Context, in *documents.DeleteDraftRequest) (*emptypb.Empty, error) {
	oid, err := cid.Decode(in.DocumentId)
	if err != nil {
		return nil, err
	}

	ocodec, ohash := ipfs.DecodeCID(oid)

	conn, release, err := api.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer release()

	if err := vcssql.DraftsDelete(conn, ohash, int(ocodec)); err != nil {
		return nil, err
	}

	if err := api.vcs.DeletePermanode(ctx, oid); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (api *Server) GetPublication(ctx context.Context, in *documents.GetPublicationRequest) (*documents.Publication, error) {
	oid, err := cid.Decode(in.DocumentId)
	if err != nil {
		return nil, err
	}

	var ver vcs.Version
	if in.Version == "" {
		ver, err = api.vcs.LoadNamedVersion(ctx, oid, api.me.AccountID(), api.me.DeviceKey().CID(), "main")
		if err != nil {
			return nil, err
		}
	} else {
		ver, err = vcs.ParseVersion(in.Version)
		if err != nil {
			return nil, err
		}
	}

	doc, err := api.getPublication(ctx, oid, ver)
	if err != nil {
		return nil, err
	}

	docpb, err := docToProto(doc)
	if err != nil {
		return nil, err
	}
	docpb.PublishTime = timestamppb.New(doc.State().UpdateTime)

	return &documents.Publication{
		Version:  ver.String(),
		Document: docpb,
		// TODO: get real latest version.
		LatestVersion: ver.String(),
	}, nil
}

func (api *Server) getPublication(ctx context.Context, oid cid.Cid, ver vcs.Version) (*vcstypes.Document, error) {
	pblk, err := vcs.LoadPermanode[vcstypes.DocumentPermanode](ctx, api.vcs.BlockGetter(), oid)
	if err != nil {
		return nil, err
	}

	p := pblk.Value

	doc := vcstypes.NewDocument(oid, p.Owner, p.CreateTime)

	if err := api.vcs.IterateChanges(ctx, oid, ver, func(c vcs.RecordedChange) error {
		var evt []vcstypes.DocumentEvent
		if err := cbornode.DecodeInto(c.Body, &evt); err != nil {
			return err
		}

		for _, e := range evt {
			if err := doc.Apply(e, c.CreateTime); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return doc, nil
}

func (api *Server) DeletePublication(ctx context.Context, in *documents.DeletePublicationRequest) (*emptypb.Empty, error) {
	c, err := cid.Decode(in.DocumentId)
	if err != nil {
		return nil, err
	}

	if err := api.vcs.DeletePermanode(ctx, c); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (api *Server) ListPublications(ctx context.Context, in *documents.ListPublicationsRequest) (*documents.ListPublicationsResponse, error) {
	conn, release, err := api.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer release()

	list, err := vcssql.PublicationsList(conn)
	if err != nil {
		return nil, err
	}

	out := &documents.ListPublicationsResponse{
		Publications: make([]*documents.Publication, len(list)),
	}

	for i, l := range list {
		aid := cid.NewCidV1(uint64(l.AccountsCodec), l.AccountsMultihash)
		pubid := cid.NewCidV1(uint64(l.ObjectsCodec), l.ObjectsMultihash).String()
		out.Publications[i] = &documents.Publication{
			Document: &documents.Document{
				Id:          pubid,
				Author:      aid.String(),
				Title:       l.PublicationsTitle,
				Subtitle:    l.PublicationsSubtitle,
				CreateTime:  &timestamppb.Timestamp{Seconds: int64(l.PublicationsCreateTime)},
				UpdateTime:  &timestamppb.Timestamp{Seconds: int64(l.PublicationsUpdateTime)},
				PublishTime: &timestamppb.Timestamp{Seconds: int64(l.PublicationsPublishTime)},
			},
			Version:       l.PublicationsLatestVersion,
			LatestVersion: l.PublicationsLatestVersion,
		}
	}

	return out, nil
}

type draft struct {
	doc       *vcstypes.Document
	wc        vcs.WorkingCopy
	oldEvents []vcstypes.DocumentEvent
}

func (api *Server) getDraft(ctx context.Context, oid cid.Cid, channel string) (*draft, error) {
	pblk, err := vcs.LoadPermanode[vcstypes.DocumentPermanode](ctx, api.vcs.BlockGetter(), oid)
	if err != nil {
		return nil, err
	}

	p := pblk.Value

	wc, err := api.vcs.LoadWorkingCopy(ctx, oid, "main")
	if err != nil {
		return nil, fmt.Errorf("failed to load working copy: %w", err)
	}

	doc := vcstypes.NewDocument(oid, p.Owner, p.CreateTime)

	if err := api.vcs.IterateChanges(ctx, oid, wc.Version(), func(c vcs.RecordedChange) error {
		var evt []vcstypes.DocumentEvent
		if err := cbornode.DecodeInto(c.Body, &evt); err != nil {
			return fmt.Errorf("failed to decode document change: %w", err)
		}

		for _, e := range evt {
			if err := doc.Apply(e, c.CreateTime); err != nil {
				return fmt.Errorf("failed to apply document event: %w", err)
			}
		}

		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed iterating changes: %w", err)
	}

	// Apply working copy events.
	var evts []vcstypes.DocumentEvent

	if wc.Data() != nil {
		if err := cbornode.DecodeInto(wc.Data(), &evts); err != nil {
			return nil, fmt.Errorf("failed to decode working copy data: %w", err)
		}
	}

	for _, e := range evts {
		if err := doc.Apply(e, wc.UpdateTime()); err != nil {
			return nil, err
		}
	}

	return &draft{
		doc:       doc,
		wc:        wc,
		oldEvents: evts,
	}, nil
}

func blockFromProto(blk *documents.Block) (vcstypes.Block, error) {
	b := vcstypes.Block{
		ID:         blk.Id,
		Type:       blk.Type,
		Attributes: blk.Attributes,
		Text:       blk.Text,
	}

	if blk.Annotations == nil {
		return b, nil
	}

	b.Annotations = make([]vcstypes.Annotation, len(blk.Annotations))

	for i, a := range blk.Annotations {
		b.Annotations[i] = vcstypes.Annotation{
			Type:       a.Type,
			Attributes: a.Attributes,
			Starts:     a.Starts,
			Ends:       a.Ends,
		}
	}

	return b, nil
}

func docToProto(d *vcstypes.Document) (*documents.Document, error) {
	docpb := &documents.Document{
		Id:         d.State().ID.String(),
		Title:      d.State().Title,
		Subtitle:   d.State().Subtitle,
		Author:     d.State().Author.String(),
		CreateTime: timestamppb.New(d.State().CreateTime),
		UpdateTime: timestamppb.New(d.State().UpdateTime), // TODO: implement real update time.
	}

	blockMap := map[string]*documents.BlockNode{}

	appendChild := func(parent string, child *documents.BlockNode) {
		if parent == crdt.RootNodeID {
			docpb.Children = append(docpb.Children, child)
			return
		}

		blk, ok := blockMap[parent]
		if !ok {
			panic("BUG: no parent " + parent + " was found yet while iterating")
		}

		blk.Children = append(blk.Children, child)
	}

	it := d.State().Tree.Iterator()

	for cur := it.NextItem(); !cur.IsZero(); cur = it.NextItem() {
		blk, ok := d.State().Blocks[cur.NodeID]
		if !ok {
			panic("BUG: node id " + cur.NodeID + " doesn't have block in the map")
		}

		child := &documents.BlockNode{Block: blockToProto(blk)}
		appendChild(cur.Parent, child)
		blockMap[cur.NodeID] = child
	}

	return docpb, nil
}

func blockToProto(blk vcstypes.Block) *documents.Block {
	bpb := &documents.Block{
		Id:         blk.ID,
		Type:       blk.Type,
		Attributes: blk.Attributes,
		Text:       blk.Text,
	}

	if blk.Annotations != nil {
		bpb.Annotations = make([]*documents.Annotation, len(blk.Annotations))
		for i, a := range blk.Annotations {
			bpb.Annotations[i] = annotationToProto(a)
		}
	}

	return bpb
}

func annotationToProto(a vcstypes.Annotation) *documents.Annotation {
	return &documents.Annotation{
		Type:       a.Type,
		Attributes: a.Attributes,
		Starts:     a.Starts,
		Ends:       a.Ends,
	}
}