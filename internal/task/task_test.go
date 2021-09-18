package task_test

import (
	"testing"

	"github.com/HALtheWise/balez/internal/task"
)

var callCount = 0

var greetT = task.Task1("greet", func(ctx *task.Context, s string) string {
	return "Hello " + embiggenT(ctx, s, 2)
})

func _embiggen(ctx *task.Context, s string, levels int) string {
	callCount += 1
	if levels == 0 {
		return s
	}
	return embiggenT(ctx, s, levels-1) + embiggenT(ctx, s, levels-1)
}

var embiggenT func(*task.Context, string, int) string

func init() {
	// Needed to break initialization cycle
	embiggenT = task.Task2("embiggen", _embiggen)
}

func TestGreet(t *testing.T) {
	ctx := task.Root()
	want := "aaaaaaaa"
	if got := embiggenT(ctx, "a", 3); got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	// callCount should have incremented
	count := callCount
	t.Logf("embiggen was called %d times", count)
	if count == 0 {
		t.Error("Counting is broken")
	}

	want = "Hello aaaa"
	if got := greetT(ctx, "a"); got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	if callCount != count {
		t.Errorf("Caching not working, calls were %d, now %d", count, callCount)
	}

	t.Logf("Statistics: %v", task.GetGlobalStats(ctx))
}
