package cache

import (
	"github.com/HALtheWise/bagel/internal/cache/graph"
)

type GlobalCache struct {
	graph.DiskCache

	// TODO(eric): These maps are temporary, the intent is to replace them with a
	// linear probing hashmap directly implemented on the capnp type.
	refsIntern    map[refKey]uint32      // Maps refs to a ref index
	funcsIntern   map[funcKey]uint32     // Maps funcs to func index
	stringsIntern map[string]StringRef   // Maps string to ref index
	funcExtraData map[uint32]interface{} // Return values from funcs that can't be serialized to disk
}

func NewGlobalCache() GlobalCache {
	return GlobalCache{
		DiskCache:     NewDiskCache(128, 128, 2048),
		refsIntern:    make(map[refKey]uint32),
		funcsIntern:   make(map[funcKey]uint32),
		stringsIntern: make(map[string]StringRef),
		funcExtraData: make(map[uint32]interface{}),
	}
}
