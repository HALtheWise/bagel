package cache

import (
	"github.com/HALtheWise/bagel/lib/cache/graph"
)

// A StringRef is a reference to an interned string
type StringRef struct {
	offset uint32
}

func (r StringRef) Pack() uint32 {
	return r.offset<<graph.RefData_bitsForKind | graph.RefData_kindRef
}

func (r StringRef) unpack(data uint32) StringRef {
	if data&KIND_MASK != graph.RefData_kindRef {
		panic("wrong kind")
	}
	return StringRef{data >> graph.RefData_bitsForKind}
}

func (r StringRef) Get(c *GlobalCache) string {
	refs, err := c.Refs()
	if err != nil {
		panic(err)
	}
	refData := refs.At(int(r.offset))

	start := refData.Left() >> graph.RefData_bitsForKind
	end := start + refData.Right()>>graph.RefData_bitsForKind

	strings, err := c.Strings()
	if err != nil {
		panic(err)
	}
	return string(strings.ToPtr().Data()[start:end])
}

func InternString(c *GlobalCache, s string) StringRef {
	if ref, ok := c.stringsIntern[s]; ok {
		return ref
	}

	start := c.StringsUsed()
	end := start + uint32(len(s))
	strings, err := c.Strings()
	if err != nil {
		panic(err)
	}

	data := strings.ToPtr().Data()
	copy(data[start:end], []byte(s))
	c.SetStringsUsed(end)

	offset := c.InternRef(
		start<<graph.RefData_bitsForKind+graph.RefData_kindString,
		uint32(len(s))<<graph.RefData_bitsForKind+graph.RefData_kindData,
	)

	ref := StringRef{offset}
	c.stringsIntern[s] = ref
	return ref
}
