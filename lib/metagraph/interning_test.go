package metagraph_test

import (
	"testing"

	"github.com/HALtheWise/bagel/lib/metagraph"
)

type X1 uint32

func TestInterning(t *testing.T) {
	var table1 metagraph.InternTable[X1, string]

	c := metagraph.NewContext()

	cat := table1.Insert(c, "cat")
	dog := table1.Insert(c, "dog")
	cat2 := table1.Insert(c, "cat")

	if cat == dog {
		t.Error("Inserted keys should not be equal")
	}

	if cat2 != cat {
		t.Error("Inserted keys should be equal")
	}

	if val := table1.Get(c, cat); val != "cat" {
		t.Errorf("Expected \"cat\", got %+v", val)
	}

	if val := table1.Get(c, dog); val != "dog" {
		t.Errorf("Expected \"dog\", got %+v", val)
	}
}
