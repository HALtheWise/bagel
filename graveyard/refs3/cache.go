package refs3

type pair struct {
	left, right Atom
}

type MemoryCache struct {
	refs     []pair
	refsDict map[pair]Atom
}

func NewCache(size int) *MemoryCache {
	return &MemoryCache{
		refs:     make([]pair, 0, size),
		refsDict: make(map[pair]Atom, size),
	}
}

func StoreRef[K, L, R IsAtom](m *MemoryCache, l L, r R) K {
	p := pair{l, r}
	if current, ok := m.refsDict[p]; ok {
		return current
	}

	k := New(KIND_REF, uint32(len(m.refs)))
	m.refsDict[p] = k

	return k
}

func RetrieveRef[K, L, R IsAtom](m *MemoryCache, k K) (L, R) {
	idx := GetRefIdx(k)
	if idx > len(m.refs) {
		panic("Index too large")
	}
	val := m.refs[idx]
	return L(val.left), R(val.right)
}
