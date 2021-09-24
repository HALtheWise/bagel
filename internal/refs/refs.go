package refs

type ref struct {
	value uint32
}

type data struct {
	left, right ref
}

type refKind uint32

const (
// KIND_STRING refKind = 1 // id is an offset into a separate strings storage system
// KIND_DATA           = 2 // id is context-specific data (probably a uint30)
// KIND_REF            = 3 // id is an offset into the refs table
)

const kindBits = 2
const kindMask = (1 << kindBits) - 1
const idBits = 32 - kindBits
const MAX_ID = (1 << idBits) - 1

func newRef(id uint32, kind refKind) ref {
	if id > MAX_ID {
		panic("ID too high")
	}
	return ref{id<<kindBits | uint32(kind)}
}

func (r ref) getKind() refKind {
	kind := r.value & kindMask
	if kind == 0 {
		panic("Invalid kind")
	}
	return refKind(kind)
}

func (r ref) getId() uint32 {
	return r.value >> kindBits
}

func (r ref) getData(state *GlobalState) data {
	if r.getKind() != KIND_REF {
		panic("Invalid ref")
	}
	return state.dataSlice[r.getId()]
}

func internData(state *GlobalState, d data) ref {
	if r := state.dataIntern[d]; r.valid() {
		return r
	}

	state.dataSlice = append(state.dataSlice, d)
	r := newRef(uint32(len(state.stringsSlice)-1), KIND_REF)

	state.dataIntern[d] = r
	return r
}

func (r ref) valid() bool {
	return r.value != 0
}

type StringRef ref

func (r StringRef) Get(state *GlobalState) string {
	if ref(r).getKind() != KIND_STRING {
		panic("Need a string")
	}
	return state.stringsSlice[ref(r).getId()]
}

func InternString(state *GlobalState, s string) StringRef {
	if r := state.stringsIntern[s]; ref(r).valid() {
		return StringRef(r)
	}

	state.stringsSlice = append(state.stringsSlice, s)
	r := newRef(uint32(len(state.stringsSlice)-1), KIND_STRING)

	state.stringsIntern[s] = r
	return StringRef(r)
}

type refLike interface {
	~struct{ value uint32 }
}

type typedRef[L, R refLike] ref

func (r typedRef[L, R]) Get(state *GlobalState) (L, R) {
	d := ref(r).getData(state)
	return L(struct{ value uint32 }(d.left)), R(struct{ value uint32 }(d.right))
}

// NOTE(eric): I couldn't figure out how to correctly ensure that the types inside TR match
// those of the member types. This function is only exported in a configured state until that gets fixed.
func intern[L refLike, R refLike, TR refLike](state *GlobalState, left L, right R) TR {
	d := data{ref(struct{ value uint32 }(left)), ref(struct{ value uint32 }(right))}
	r := internData(state, d)

	return TR(struct{ value uint32 }(r))
}

type PkgRef = typedRef[StringRef, StringRef]

var InternPkg = intern[StringRef, StringRef, PkgRef]

type LabelRef = typedRef[PkgRef, StringRef]

var InternLabel = intern[PkgRef, StringRef, LabelRef]

// var _ PkgRef = Intern(&GlobalState{}, StringRef{}, StringRef{})

// type PackageData data
// func (r PkgRef) Get(state *GlobalState) PackageData {
// 	return PackageData(ref(r).getData(state))
// }
// func (d PackageData) Workspace() StringRef {
// 	return StringRef(d.left)
// }
// func (d PackageData) Path() StringRef {
// 	return StringRef(d.right)
// }

type GlobalState struct {
	dataSlice     []data
	dataIntern    map[data]ref
	stringsSlice  []string
	stringsIntern map[string]ref
}
