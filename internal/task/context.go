package task

import "fmt"

// TODO(HALtheWise): Include a context.Context in this
// TODO(HALtheWise): Consider using automatically assigned IDs instead of strings here
type Context struct {
	name   string
	global *globalContext
}

type globalContext struct {
	// Mapping from ID to typed memo table
	memoTables []interface{}
	stats      GlobalStats
}

type GlobalStats struct {
	cacheHits, cacheMisses int
}

func GetGlobalStats(c *Context) *GlobalStats {
	return &c.global.stats
}

func (g *GlobalStats) String() string {
	return fmt.Sprintf("Cache hits: %d ; Cache Misses: %d", g.cacheHits, g.cacheMisses)
}

func Root() *Context {
	return &Context{
		name: "",
		global: &globalContext{
			memoTables: make([]interface{}, getMaxID()+1),
		},
	}
}

func getTypedMemo[Arg comparable, Result any](g *globalContext, id int) map[Arg]Result {
	if id >= len(g.memoTables) {
		newTable := make([]interface{}, id+1)
		copy(newTable, g.memoTables)
		g.memoTables = newTable
	}
	if memo := g.memoTables[id]; memo != nil {
		if typedMemo, ok := memo.(map[Arg]Result); ok {
			return typedMemo
		} else {
			panic(fmt.Errorf("ID %d reused by multiple tasks with different types, %T, %T", id, typedMemo, memo))
		}
	} else {
		typedMemo := make(map[Arg]Result)
		g.memoTables[id] = typedMemo
		return typedMemo
	}
}
