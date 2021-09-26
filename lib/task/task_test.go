package task_test

import (
	"testing"

	"github.com/HALtheWise/bagel/lib/task"
)

var callCount = 0

var T_greet = task.Task1("greet", func(ctx *task.Context, s string) string {
	return "Hello " + T_embiggen(ctx, s, 2)
})

func _embiggen(ctx *task.Context, s string, levels int) string {
	callCount += 1
	if levels == 0 {
		return s
	}
	return T_embiggen(ctx, s, levels-1) + T_embiggen(ctx, s, levels-1)
}

var T_embiggen func(*task.Context, string, int) string

func init() {
	// Needed to break initialization cycle
	T_embiggen = task.Task2("embiggen", _embiggen)
}

func TestGreet(t *testing.T) {
	ctx := task.Root()
	want := "aaaaaaaa"
	if got := T_embiggen(ctx, "a", 3); got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	// callCount should have incremented
	count := callCount
	t.Logf("embiggen was called %d times", count)
	if count == 0 {
		t.Error("Counting is broken")
	}

	want = "Hello aaaa"
	if got := T_greet(ctx, "a"); got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	if callCount != count {
		t.Errorf("Caching not working, calls were %d, now %d", count, callCount)
	}

	t.Logf("Statistics: %v", task.GetGlobalStats(ctx))
}
