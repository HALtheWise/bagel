package task2

// an anyRef has low bits of metadata and high bits of an index.
type anyRef uint32

type refKind uint32

const DEBUG_ASSERTS bool = true

const kindBits = 2
const dataBits = 32 - kindBits

const kind_mask = 1<<kindBits - 1
const data_mask = ((1 << dataBits) - 1) << kindBits

const INVALID_REF = anyRef(0)

const (
	KIND_INVALID refKind = 0
	KIND_STRING  refKind = 1
	KIND_STRUCT  refKind = 2
	KIND_DATA    refKind = 3
)

type StringRef struct {
	ref anyRef
}

func (s StringRef) Lookup(c *Context) string {
	if DEBUG_ASSERTS && s.ref == INVALID_REF {
		panic(0)
	}
	return c.Global.strings[s.ref.Data()]
}

func InternString(c *Context, s string) StringRef {
	if existing, ok := c.Global.stringsMap[s]; ok {
		return StringRef{existing}
	}

	ref := newRef(KIND_STRING, uint32(len(c.Global.strings)))
	c.Global.strings = append(c.Global.strings, s)
	c.Global.stringsMap[s] = ref

	return StringRef{ref}
}

type DataRef struct {
	ref anyRef
}

func (d DataRef) Get() uint32 {
	return d.Get()
}

func InternData(data uint32) DataRef {
	if data >= 1<<dataBits {
		panic("too big")
	}
	return DataRef{newRef(KIND_DATA, data)}
}

func (a anyRef) Kind() refKind {
	return refKind(kindBits & a)
}

func (a anyRef) Data() uint32 {
	return uint32((a & data_mask) >> kindBits)
}

func newRef(kind refKind, data uint32) anyRef {
	if DEBUG_ASSERTS && (kind >= 1<<kindBits || kind == 0) {
		panic(0)
	}
	return anyRef(data<<kindBits | uint32(kind))
}
