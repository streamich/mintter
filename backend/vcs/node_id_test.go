package vcs

import (
	"io"
	"math"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"

	cbornode "github.com/ipfs/go-ipld-cbor"
	"github.com/stretchr/testify/require"
)

func TestNodeID2String(t *testing.T) {
	require.Equal(t, "b1", NodeIDFromString("b1").String())
	require.Equal(t, "ab1", NodeIDFromString("ab1").String())
	require.Equal(t, "$ROOT", RootNode.String())
	require.Equal(t, "$TRASH", TrashNode.String())

	b := [8]byte{234, 255, 135, 120, 50, 60, 70, 0}
	nid := *(*NodeID)(unsafe.Pointer(&b))
	require.Equal(t, "6v+HeDI8Rg", nid.String())

	require.Equal(t, TrashNode, NodeIDFromString("$TRASH"))
	require.Equal(t, RootNode, NodeIDFromString("$ROOT"))
}

func TestIncrementCounter(t *testing.T) {
	old := uint14(atomic.LoadUint64(&nidCounter))

	for i := 0; i < int(maxUint14)*3; i++ {
		c := nidCounterInc()

		if c > maxUint14 {
			t.Fatalf("%d must not be larger than %d", c, maxUint14)
		}

		diff := int64(c) - int64(old)
		if diff != 1 && diff != -1 && diff != -int64(maxUint14) {
			t.Fatalf("old=%d, new=%d, diff=%d", old, c, diff)
		}
		old = c
	}
}

func TestNodeIDv1BitPacking(t *testing.T) {
	nid := newNodeIDv1(time.Unix(1666354607, 0), 2, 2)
	want := NodeID(447308658787778562)
	require.Equal(t, want, nid)
}

func TestRandomUint14(t *testing.T) {
	for i := 0; i < 10000; i++ {
		require.True(t, randomUint14() <= maxUint14)
	}
}

func TestUint14(t *testing.T) {
	require.Equal(t, maxUint14, uint14(math.MaxUint64))
}

func TestNodeIDDifferent(t *testing.T) {
	ids := make(map[NodeID]int, 1500)

	for i := 0; i < 1500; i++ {
		ids[NewNodeIDv1(time.Now())]++
	}

	for id, count := range ids {
		if count > 1 {
			t.Fatalf("%d was generated %d times", id, count)
		}
	}
}

type Thing struct {
	ID   string
	Body Body
}

type Body struct {
	Name string
}

func (b Body) MarshalCBOR(w io.Writer) error {
	data, err := cbornode.DumpObject([]any{b.Name})
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

func (b Body) EncodeDagCbor(w io.Writer) error {
	data, err := cbornode.DumpObject([]any{b.Name})
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}
