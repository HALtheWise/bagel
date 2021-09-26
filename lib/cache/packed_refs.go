package cache

import (
	"github.com/HALtheWise/bagel/lib/cache/graph"
)

const (
	MAX_OFFSET        = 1<<(32-graph.RefData_bitsForKind) - 1
	KIND_MASK  uint32 = 1<<graph.RefData_bitsForKind - 1
)

type Ref[T Ref[T]] interface {
	Pack() uint32
	unpack(uint32) T
}

func fromPacked3[R Ref[R]](v uint32) R {
	var zero R
	return zero.unpack(v)
}

// A genericRef helps implement useful helper functions for a Ref stored in the Refs array of the global context.
type genRef3[L Ref[L], R Ref[R]] struct {
	// type genRef3[left any, right any, self is[genRef3[left, right, self]]] struct {
	offset uint32
}

func (r genRef3[L, R]) Get(c *GlobalCache) (L, R) {
	refs, err := c.Refs()
	if err != nil {
		panic(err)
	}
	refData := refs.At(int(r.offset))

	return fromPacked3[L](refData.Left()),
		fromPacked3[R](refData.Right())
}

func (r *genRef3[L, R]) fillFromIntern(c *GlobalCache, left L, right R) {
	r.offset = c.InternRef(left.Pack(), right.Pack())
}

func (r *genRef3[_, _]) fillFromUnpack(data uint32) {
	if data&KIND_MASK != graph.RefData_kindRef {
		panic("Wrong kind")
	}
	r.offset = data >> graph.RefData_bitsForKind
}

func (r genRef3[_, _]) Pack() uint32 {
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

func InternPackage3(c *GlobalCache, workspace, path PackageRef3) PackageRef3 {
	var result PackageRef3
	result.fillFromIntern(c, workspace, path)
	return result
}
