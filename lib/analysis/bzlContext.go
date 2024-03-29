package analysis

import (
	"go.starlark.net/starlark"

	"github.com/HALtheWise/bagel/lib/core"
	"github.com/HALtheWise/bagel/lib/refs"
)

var _ starlark.HasAttrs = &BzlLabel{}

// https://docs.bazel.build/versions/main/skylark/lib/Label.html
type BzlLabel struct {
	ctx    *core.Context
	label  refs.LabelRef
	frozen bool
}

func (l *BzlLabel) AttrNames() []string {
	return []string{"name", "package", "workspace_name"} // TODO: relative, workspace_root
}

func (l *BzlLabel) Attr(name string) (starlark.Value, error) {
	c := l.ctx
	switch name {
	case "name":
		return starlark.String(l.label.Get(c).Name.Get(c)), nil
	case "package":
		return starlark.String(l.label.Get(c).Pkg.Get(c).RelPath.Get(c)), nil
	case "workspace_name":
		return starlark.String(l.label.Get(c).Pkg.Get(c).Workspace.Get(c)), nil
	}
	return nil, nil
}

func (l *BzlLabel) String() string        { return l.label.String() }
func (l *BzlLabel) Type() string          { return "label" }
func (l *BzlLabel) Freeze()               { l.frozen = true }
func (l *BzlLabel) Truth() starlark.Bool  { return true }
func (l *BzlLabel) Hash() (uint32, error) { return starlark.String(l.String()).Hash() }

// https://docs.bazel.build/versions/main/skylark/lib/ctx.html
var _ starlark.HasAttrs = &BzlLabel{}

type BzlCtx struct {
	ctx            *core.Context
	clabel         refs.CLabelRef
	pkg            refs.PackageRef
	actions        []*Action
	files_declared []BzlFile
	attrs          map[string]AnalyzedAttrValue
}

func (c *BzlCtx) AttrNames() []string {
	return []string{"label", "actions", "attr"} // TODO: lots
}

func (c *BzlCtx) Attr(name string) (starlark.Value, error) {
	switch name {
	case "label":
		return &BzlLabel{label: c.clabel.Get(c.ctx).Label, frozen: false, ctx: c.ctx}, nil
	case "actions":
		return &BzlActions{c}, nil
	case "attr":
		return &BzlAttrGetter{c, ATTR_GETTER}, nil
	case "file":
		return &BzlAttrGetter{c, FILE_GETTER}, nil
	}
	return nil, nil
}

func (c *BzlCtx) String() string        { return "ctx for " + c.clabel.String() }
func (c *BzlCtx) Type() string          { return "ctx" }
func (c *BzlCtx) Freeze()               { panic("Cannot freeze BzlCtx") }
func (c *BzlCtx) Truth() starlark.Bool  { return starlark.True }
func (c *BzlCtx) Hash() (uint32, error) { panic("not implemented") }
