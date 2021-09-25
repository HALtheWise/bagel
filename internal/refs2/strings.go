package refs2

import "github.com/HALtheWise/bagel/internal/dcache/graph"

// A StringRef is a reference to an interned string
type StringRef struct {
	offset uint32
}

func (r StringRef) pack() uint32 {
	return r.offset<<graph.RefData_bitsForKind | graph.RefData_kindRef
}

func (r StringRef) unpack(data uint32) StringRef {
	if data&KIND_MASK != graph.RefData_kindRef {
		panic("wrong kind")
	}
	return StringRef{data >> graph.RefData_bitsForKind}
}

func (r StringRef) Get(c *GlobalContext) string {
	refs, err := c.Cache.Refs()
	if err != nil {
		panic(err)
	}
	refData := refs.At(int(r.offset))

	start := refData.Left() >> graph.RefData_bitsForKind
	end := start + refData.Right()>>graph.RefData_bitsForKind

	strings, err := c.Cache.Strings()
	if err != nil {
		panic(err)
	}
	return string(strings.ToPtr().Data()[start:end])
}

func InternString(c *GlobalContext, s string) StringRef {
	if ref, ok := c.stringsIntern[s]; ok {
		return ref
	}

	start := c.Cache.StringsUsed()
	end := start + uint32(len(s))
	strings, err := c.Cache.Strings()
	if err != nil {
		panic(err)
	}
	copy([]byte(s), string(strings.ToPtr().Data())[start:end])

	offset := c.Cache.InternRef(
		start<<graph.RefData_bitsForKind+graph.RefData_kindString,
		uint32(len(s))<<graph.RefData_bitsForKind+graph.RefData_kindData,
	)

	return StringRef{offset}
}
