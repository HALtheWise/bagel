package task

import (
	"fmt"

	"github.com/HALtheWise/bagel/lib/cache"
)

// TODO(HALtheWise): Include a context.Context in this
// TODO(HALtheWise): Consider using automatically assigned IDs instead of strings here
type Context struct {
	name   string
	Global *globalContext
}

type globalContext struct {
	// Mapping from ID to typed memo table
	Cache cache.GlobalCache
	stats Stats
}

type Stats struct {
	cacheHits, cacheMisses int
}

func GetGlobalStats(c *Context) *Stats {
	return &c.Global.stats
}

func (g *Stats) String() string {
	return fmt.Sprintf("Cache hits: %d ; Cache Misses: %d", g.cacheHits, g.cacheMisses)
}

func Root() *Context {
	return &Context{
		name: "",
		Global: &globalContext{
			Cache: cache.NewGlobalCache(),
		},
	}
}
