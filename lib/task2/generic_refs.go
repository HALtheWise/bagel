package task2

type GenericRef[T GenericRef[T]] interface {
	toRef() anyRef
	fromRef(anyRef) T
}

type CompositeRef[L GenericRef[L], R GenericRef[R]] struct {
	ref anyRef
}

func (r CompositeRef[L, R]) toRef() anyRef {
	return r.ref
}

func (r CompositeRef[L, R]) fromRef(ref anyRef) CompositeRef[L, R] {
	return CompositeRef[L, R]{ref}
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
