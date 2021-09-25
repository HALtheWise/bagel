package refs

import (
	"github.com/HALtheWise/bagel/internal/dcache"
	"github.com/HALtheWise/bagel/internal/dcache/graph"
)

type GlobalContext struct {
	Cache dcache.DCache

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
