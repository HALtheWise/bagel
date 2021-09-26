package cache

import (
	"os/exec"

	capnp "zombiezen.com/go/capnproto2"

	"github.com/HALtheWise/bagel/internal/cache/graph"
)

var version string

type refKey struct {
	left, right uint32
}

type funcKey struct {
	kind, arg uint32
}

// The Intern functions return the offset in the Addr or Funcs array containing a reference to the provided object.
// If the object was already in the array, it will not be modified.
func (d *GlobalCache) InternRef(left, right uint32) uint32 {
	key := refKey{left, right}
	if addr, ok := d.refsIntern[key]; ok {
		return addr
	}

	refs, err := d.Refs()
	if err != nil {
		panic(err)
	}
	addr := uint32(len(d.refsIntern))
	ref := refs.At(int(addr))

	ref.SetLeft(left)
	ref.SetRight(right)

	d.refsIntern[key] = addr
	return addr
}

func (d *GlobalCache) InternFunc(kind, arg uint32) uint32 {
	key := funcKey{kind, arg}
	if addr, ok := d.funcsIntern[key]; ok {
		return addr
	}

	funcs, err := d.Funcs()
	if err != nil {
		panic(err)
	}

	addr := uint32(len(d.funcsIntern))
	funcs.At(int(addr)).SetKind(kind)
	funcs.At(int(addr)).SetArg(arg)
	d.funcsIntern[key] = addr
	return addr
}

func NewDiskCache(refsSize, funcsSize, stringsSize int32) graph.DiskCache {
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
