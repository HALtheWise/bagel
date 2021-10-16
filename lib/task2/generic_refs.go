package task2

type baseRef = struct{ ref anyRef }

type GenericRef interface {
	~baseRef
}

type CompositeRef[L, R GenericRef] struct {
	ref anyRef
}

func (ref CompositeRef[L, R]) Lookup(c *Context) (L, R) {
	s := c.Global.structs[ref.ref.Data()]
	baseLeft := baseRef{s.left}
	baseRight := baseRef{s.right}
	return L(baseLeft), R(baseRight)
}

func InternCompositeRef[L, R GenericRef](c *Context, l L, r R) CompositeRef[L, R] {
	baseLeft := baseRef(l)
	baseRight := baseRef(r)

	s := anyStruct{baseLeft.ref, baseRight.ref}
	c.Global.structs = append(c.Global.structs, s)

	ref := newRef(KIND_STRUCT, uint32(len(c.Global.structs)-1))
	c.Global.structsMap[s] = ref

	return CompositeRef[L, R](baseRef{ref})
}
