package dcache

import (
	"os/exec"

	"github.com/HALtheWise/bales/internal/dcache/graph"
	capnp "zombiezen.com/go/capnproto2"
)

var version string

func New(refsSize, funcsSize, stringsSize int32) graph.DiskCache {
	// Make a brand new empty message.  A Message allocates Cap'n Proto structs.
	_, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic(err)
	}

	// Create a root struct.  Every message must have a root struct.
	cache, err := graph.NewDiskCache(seg)
	if err != nil {
		panic(err)
	}

	if version == "" {
		versionBytes, err := exec.Command("git", "rev-parse", "HEAD").Output()
		if err != nil {
			panic(err)
		}
		version = string(versionBytes)
	}

	cache.SetVersion(version)

	// Create refs table
	refData, err := graph.NewRefData_List(seg, refsSize)
	if err != nil {
		panic(err)
	}
	err = cache.SetRefs(refData)
	if err != nil {
		panic(err)
	}

	// Create funcs table
	funcsData, err := graph.NewFuncObj_List(seg, funcsSize)
	if err != nil {
		panic(err)
	}
	err = cache.SetFuncs(funcsData)
	if err != nil {
		panic(err)
	}

	// Create strings data
	stringsData, err := capnp.NewInt8List(seg, stringsSize)
	if err != nil {
		panic(err)
	}
	err = cache.SetStrings(capnp.UInt8List(stringsData))
	if err != nil {
		panic(err)
	}

	return cache
}
