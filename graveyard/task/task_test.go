package task_test

import (
	"fmt"
	"testing"

	capnp "zombiezen.com/go/capnproto2"

	"github.com/HALtheWise/bagel/lib/cache"
	"github.com/HALtheWise/bagel/lib/cache/graph"
	"github.com/HALtheWise/bagel/lib/task"
)

var callCount = 0

var T_greet func(ctx *task.Context, s cache.StringRef) string

func init() {
	T_greet = task.GoTask("greet", func(ctx *task.Context, s cache.StringRef) string {
		str := s.Get(&ctx.Global.Cache)
		if len(str) > 20 {
			return str
		}
		empty := cache.InternString(&ctx.Global.Cache, "")
		code := T_embiggen(ctx, empty)
		return T_greet(ctx, cache.InternString(&ctx.Global.Cache, fmt.Sprintf("hello %s %d", str, code.Left())))
	})
}

var T_embiggen = task.DiskTask("embiggen", func(c *task.Context, s cache.StringRef, segment *capnp.Segment) graph.RefData {
	d, err := graph.NewRefData(segment)
	if err != nil {
		panic(err)
	}
	d.SetLeft(42)
	return d
})

func TestGreet(t *testing.T) {
	ctx := task.Root()

	want := "hello hello world 42 42"
	if got := T_greet(ctx, cache.InternString(&ctx.Global.Cache, "world")); got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	t.Logf("Statistics: %v", task.GetGlobalStats(ctx))
}
