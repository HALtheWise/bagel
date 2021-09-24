package refs2

import (
	"github.com/HALtheWise/bagel/internal/dcache/graph"
)

// A genericRef helps implement useful helper functions for a Ref stored in the Refs array of the global context.
type genRef3[L packer[L], R packer[R]] struct {
	// type genRef3[left any, right any, self is[genRef3[left, right, self]]] struct {
	offset uint32
}

func (r genRef3[L, R]) Get(c *GlobalContext) (L, R) {
	refs, err := c.Cache.Refs()
	if err != nil {
		panic(err)
	}
	refData := refs.At(int(r.offset))

	return fromPacked3[L](refData.Left()),
		fromPacked3[R](refData.Right())
}

func (r *genRef3[L, R]) fillFromIntern(c *GlobalContext, left L, right R) {
	r.offset = c.Cache.InternRef(left.pack(), right.pack())
}

func (r *genRef3[_, _]) fillFromUnpack(data uint32) {
	if data&KIND_MASK != graph.RefData_kindRef {
		panic("Wrong kind")
	}
	r.offset = data >> graph.RefData_bitsForKind
}

func (r genRef3[_, _]) pack() uint32 {
	if r.offset > MAX_OFFSET {
		panic("Offset too large")
	}
	return r.offset<<graph.RefData_bitsForKind | graph.RefData_kindRef
}

type PackageRef3 struct {
	genRef3[PackageRef3, PackageRef3]
}

func (PackageRef3) unpack(data uint32) PackageRef3 {
	var result PackageRef3
	result.fillFromUnpack(data)
	return result
}

func InternPackage3(c *GlobalContext, workspace, path PackageRef3) PackageRef3 {
	var result PackageRef3
	result.fillFromIntern(c, workspace, path)
	return result
}

func xtest() {
	r := InternPackage3(&GlobalContext{}, PackageRef3{}, PackageRef3{})
	r.Get(&GlobalContext{})

	// fromPacked[LabelRef](0)
}
