package task2_test

import (
	"testing"

	"github.com/HALtheWise/bagel/lib/task2"
)

type PairRef struct {
	task2.CompositeRef[task2.StringRef, task2.StringRef]
}

func InternPair(g *task2.Context, left, right task2.StringRef) PairRef {
	return PairRef{task2.InternCompositeRef(g, left, right)}
}

type MetaPair struct {
	task2.CompositeRef[PairRef, task2.DataRef]
}

func TestPairRef(t *testing.T) {
	c := &task2.Context{
		TaskID: 0,
		Global: task2.NewGlobalContext(),
	}
	pair := InternPair(c, task2.InternString(c, "a"), task2.InternString(c, "b"))

	l, r := pair.Lookup(c)

	t.Logf("%v %v", l.Lookup(c), r.Lookup(c))
}
