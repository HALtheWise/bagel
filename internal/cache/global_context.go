package cache

import (
	"github.com/HALtheWise/bagel/internal/cache/graph"
)

type GlobalContext struct {
	graph.DiskCache

	// TODO(eric): These maps are temporary, the intent is to replace them with a
	// linear probing hashmap directly implemented on the capnp type.
	refsIntern  map[refKey]uint32  // Maps refs to a ref index
	funcsIntern map[funcKey]uint32 // Maps funcs to func index

	// Temporary until I get an in-place hash table working
	stringsIntern map[string]StringRef
}

type packer[T packer[T]] interface {
	pack() uint32
	unpack(uint32) T
}

const (
	MAX_OFFSET        = 1<<(32-graph.RefData_bitsForKind) - 1
	KIND_MASK  uint32 = 1<<graph.RefData_bitsForKind - 1
)

func fromPacked3[R packer[R]](v uint32) R {
	var zero R
	return zero.unpack(v)
}
