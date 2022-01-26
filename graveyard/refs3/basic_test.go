package refs3_test

import (
	"testing"

	"github.com/HALtheWise/bagel/lib/refs3"
)

type Int refs3.Atom

func (i Int) Value() int {
	return int(refs3.GetData(i)) - 1<<29
}

func MakeInt(value int) Int {
	return refs3.New(refs3.KIND_DATA, uint32(value+1<<29))
}

func TestBasic(t *testing.T) {
	i := MakeInt(42)
	if i.Value() != 42 {
		t.Error()
	}
}
