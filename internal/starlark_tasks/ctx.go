package starlark_tasks

import (
	"fmt"

	"github.com/HALtheWise/balez/internal/labels"
	"go.starlark.net/starlark"
)

type BzlLabel struct {
	label  labels.Label
	frozen bool
}

func (l *BzlLabel) AttrNames() []string {
	return []string{"name", "package", "workspace_name"} // TODO: relative, workspace_root
}

func (l *BzlLabel) Attr(name string) (starlark.Value, error) {
	switch name {
	case "name":
		return starlark.String(l.label.Name), nil
	case "package":
		return starlark.String(l.label.Pkg), nil
	case "workspace_name":
		return starlark.String(l.label.Repo), nil
	}
	return starlark.None, fmt.Errorf("Unimplemented label.%s", name)
}

func (l *BzlLabel) String() string        { return l.label.String() }
func (l *BzlLabel) Type() string          { return "label" }
func (l *BzlLabel) Freeze()               { l.frozen = true }
func (l *BzlLabel) Truth() starlark.Bool  { return true }
func (l *BzlLabel) Hash() (uint32, error) { return starlark.String(l.String()).Hash() }

// https://docs.bazel.build/versions/main/skylark/lib/ctx.html
type BzlCtx struct {
	label BzlLabel
}

func (c *BzlCtx) AttrNames() []string {
	return []string{"label"} // TODO: lots
}

func (c *BzlCtx) Attr(name string) (starlark.Value, error) {
	switch name {
	case "label":
		return &c.label, nil
	}
	return starlark.None, fmt.Errorf("Unimplemened ctx.%s", name)
}

func (c *BzlCtx) String() string        { return "ctx for " + c.label.String() }
func (c *BzlCtx) Type() string          { return "ctx" }
func (c *BzlCtx) Freeze()               { panic("Cannot freeze BzlCtx") }
func (c *BzlCtx) Truth() starlark.Bool  { return starlark.True }
func (c *BzlCtx) Hash() (uint32, error) { panic("not implemented") }
