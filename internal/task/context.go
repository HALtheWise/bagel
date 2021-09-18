package task

import "fmt"

// TODO(HALtheWise): Include a context.Context in this
// TODO(HALtheWise): Consider using automatically assigned IDs instead of strings here
type Context struct {
	name   string
	global *globalContext
}

type globalContext struct {
	memoTables map[string]interface{}
	Stats      GlobalStats
}

type GlobalStats struct {
	cacheHits, cacheMisses int
}

func GetGlobalStats(c *Context) *GlobalStats {
	return &c.global.Stats
}

func (g *GlobalStats) String() string {
	return fmt.Sprintf("Cache hits: %d ; Cache Misses: %d", g.cacheHits, g.cacheMisses)
}

func Root() *Context {
	return &Context{
		name: "",
		global: &globalContext{
			memoTables: make(map[string]interface{}),
		},
	}
}

func getTypedMemo[Arg comparable, Result any](g *globalContext, name string) map[Arg]Result {
	if memo, ok := g.memoTables[name]; ok {
		if typedMemo, ok := memo.(map[Arg]Result); ok {
			return typedMemo
		} else {
			panic(fmt.Errorf("Name %s reused by multiple tasks with different types", name))
		}
	} else {
		typedMemo := make(map[Arg]Result)
		g.memoTables[name] = typedMemo
		return typedMemo
	}
}
