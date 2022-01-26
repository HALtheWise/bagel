package core_test

import (
	"fmt"
	"testing"

	"github.com/HALtheWise/bagel/lib/core"
)

type StringRef uint32

var StringInterner core.InternTable[StringRef, string]

func TestHello(t *testing.T) {
	executions := 0

	hello := core.MemoFunc1("hello",
		func(c *core.Context, name StringRef) string {
			executions++
			return fmt.Sprintf("Hello %s", StringInterner.Get(c, name))
		})

	c := core.NewContext()
	glados := StringInterner.Insert(c, "GLaDOS")

	if result := hello(c, glados); result != "Hello GLaDOS" {
		t.Error("Expected Hello GLaDOS, got", result)
	}

	hello(c, glados)
	hello(c, glados)
	hello(c, glados)

	if executions != 1 {
		t.Error("Func executed too many times: ", executions)
	}

	cortana := StringInterner.Insert(c, "cortana")

	hello(c, cortana)
	hello(c, cortana)
	if executions != 2 {
		t.Error("Func executed too many times: ", executions)
	}
}

func TestFib(t *testing.T) {
	executions := 0

	var fib func(*core.Context, uint32) int
	fib = core.MemoFunc1("fib", func(c *core.Context, a uint32) int {
		executions++
		if a <= 2 {
			return 1
		}
		return fib(c, a-1) + fib(c, a-2)
	})

	c := core.NewContext()

	for i := 1; i <= 20; i++ {
		t.Log(i, fib(c, uint32(i)))
	}

	if executions != 20 {
		t.Error("Function executed too many times", executions)
	}
}

func TestExpand(t *testing.T) {
	executions := 0

	var expand func(*core.Context, StringRef, uint32) string
	expand = core.MemoFunc2("expand",
		func(c *core.Context, name StringRef, depth uint32) string {
			executions++
			if depth == 0 {
				return StringInterner.Get(c, name)
			}
			return expand(c, name, depth-1) + expand(c, name, depth-1)
		})

	c := core.NewContext()

	eric := StringInterner.Insert(c, "eric")
	if expand(c, eric, 2) != "ericericericeric" {
		t.Error("Expand wrong")
	}
	if executions != 3 {
		t.Error("caching broken")
	}
}
