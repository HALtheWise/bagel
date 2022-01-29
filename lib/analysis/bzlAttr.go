package analysis

import (
	"sort"

	"go.starlark.net/starlark"

	"github.com/HALtheWise/bagel/lib/loading"
)

// This is ctx.attr
type BzlAttr struct {
	ctx *BzlCtx
}

var _ starlark.HasAttrs = &BzlAttr{}

func (a BzlAttr) AttrNames() []string {
	names := make([]string, 0, len(a.ctx.attrs))
	for name := range a.ctx.attrs {
		names = append(names, name)
	}

	sort.Strings(names)
	return names
}

func (a BzlAttr) Attr(name string) (starlark.Value, error) {
	attr, ok := a.ctx.attrs[name]
	if !ok {
		return nil, nil
	}

	switch attr.Kind {
	case loading.STRING_ATTR:
		return starlark.String(attr.StringValue), nil
	default:
		panic("Unknown attr kind")
	}
}

func (a BzlAttr) String() string        { return "attr getters for " + a.ctx.clabel.String() }
func (a BzlAttr) Type() string          { return "attr" }
func (a BzlAttr) Freeze()               { a.ctx.Freeze() }
func (a BzlAttr) Truth() starlark.Bool  { return starlark.True }
func (a BzlAttr) Hash() (uint32, error) { panic("not implemented") }
