package task

import (
	"fmt"

	"github.com/HALtheWise/bagel/internal/cache"
)

// TODO(HALtheWise): Include a context.Context in this
// TODO(HALtheWise): Consider using automatically assigned IDs instead of strings here
type Context struct {
	name   string
	global *globalContext
}

type globalContext struct {
	// Mapping from ID to typed memo table
	cache cache.GlobalCache
	stats Stats
}

type Stats struct {
	cacheHits, cacheMisses int
}

func GetGlobalStats(c *Context) *Stats {
	return &c.global.stats
}

func (g *Stats) String() string {
	return fmt.Sprintf("Cache hits: %d ; Cache Misses: %d", g.cacheHits, g.cacheMisses)
}

func Root() *Context {
	return &Context{
		name: "",
		global: &globalContext{
			cache: cache.NewGlobalCache(),
		},
	}
}

func getTypedMemo[Arg comparable, Result any](g *globalContext, id int) map[Arg]Result {
	if id >= len(g.cache) {
		newTable := make([]interface{}, id+1)
		copy(newTable, g.cache)
		g.cache = newTable
	}
	if memo := g.cache[id]; memo != nil {
		if typedMemo, ok := memo.(map[Arg]Result); ok {
			return typedMemo
		} else {
			panic(fmt.Errorf("ID %d reused by multiple tasks with different types, %T, %T", id, typedMemo, memo))
		}
	} else {
		typedMemo := make(map[Arg]Result)
		g.cache[id] = typedMemo
		return typedMemo
	}
}
