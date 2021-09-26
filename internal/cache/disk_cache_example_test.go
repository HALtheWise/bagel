package cache_test

import (
	"os"
	"testing"

	"github.com/HALtheWise/bagel/internal/cache/graph"
	capnp "zombiezen.com/go/capnproto2"
)

func TestWrite(t *testing.T) {
	version := "dev"
	size := int32(4)

	// Make a brand new empty message.  A Message allocates Cap'n Proto structs.
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic(err)
	}

	// Create a root struct.  Every message must have a root struct.
	cache, err := graph.NewDiskCache(seg)
	if err != nil {
		panic(err)
	}
	cache.SetVersion(version)

	// Create refs table
	refData, err := graph.NewRefData_List(seg, size)
	if err != nil {
		panic(err)
	}
	err = cache.SetRefs(refData)
	if err != nil {
		panic(err)
	}

	// Create funcs table
	funcsData, err := graph.NewFuncObj_List(seg, size)
	if err != nil {
		panic(err)
	}
	err = cache.SetFuncs(funcsData)
	if err != nil {
		panic(err)
	}

	refData.At(2).SetLeft(42)
	refData.At(2).SetRight(43)
	refData.At(3).SetLeft(44)
	refData.At(3).SetRight(45)

	t.Log(len(seg.Data()), seg.Data())
	// Write the message to stdout.
	err = capnp.NewEncoder(os.Stdout).Encode(msg)
	if err != nil {
		panic(err)
	}
}
