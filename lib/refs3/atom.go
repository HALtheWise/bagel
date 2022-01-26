package refs3

const kindBits = 2
const kindMask = 1<<kindBits - 1

type AtomKind uint32

const (
	KIND_INVALID AtomKind = 0
	KIND_REF     AtomKind = 1
	KIND_STRING  AtomKind = 2
	KIND_DATA    AtomKind = 3
)

// A Atom is a 32 bit value holding an object reference that can be used as an argument
// type for a memoized function. Kind() returns the type of value it holds:
// - KIND_DATA: This atom holds a uint30 (probably a length or index type)
// - KIND_STRING: This atom refers to a location in the strings interning array
// - KIND_REF: This atom refers to an interned "object" consisting of two other atoms.
//
// Atom is not intended to be used directly, instead a instance should be defined.
type Atom = struct {
	packed uint32
}

type IsAtom interface {
	~Atom
}

func Kind[R IsAtom](atom R) AtomKind {
	return AtomKind(Atom(atom).packed & kindMask)
}

func New[R IsAtom](kind AtomKind, value uint32) R {
	if value >= 1<<(32-kindBits) {
		panic("Value too large")
	}
	if kind > KIND_DATA || kind == KIND_INVALID {
		panic("Invalid kind")
	}

	return R(Atom{value<<kindBits | uint32(kind)})
}

func GetData[R IsAtom](atom R) uint32 {
	if Kind(atom) != KIND_DATA {
		panic("Atom was not data")
	}
	return Atom(atom).packed >> kindBits
}

func GetRefIdx[R IsAtom](atom R) int {
	if Kind(atom) != KIND_REF {
		panic("Atom was not ref")
	}
	return int(Atom(atom).packed >> kindBits)
}

func GetStringOffset[R IsAtom](atom R) int {
	if Kind(atom) != KIND_STRING {
		panic("Atom was not string")
	}
	return int(Atom(atom).packed >> kindBits)
}
