package analysis

import (
	"fmt"

	"go.starlark.net/starlark"

	"github.com/HALtheWise/bagel/lib/core"
	"github.com/HALtheWise/bagel/lib/refs"
)

// https://docs.bazel.build/versions/main/skylark/lib/actions.html#declare_file
func actionDeclareFile(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var filename string
	if err := starlark.UnpackArgs("declare_file", args, kwargs, "filename", &filename); err != nil {
		return nil, err
	}

	actions := fn.Receiver().(BzlActions)
	c := actions.ctx.ctx
	label := refs.Label{
		Pkg:  actions.ctx.pkg,
		Name: refs.StringTable.Insert(c, filename)}

	cfile := refs.CFile{
		Location: refs.LabelTable.Insert(c, label),
		Source:   actions.ctx.clabel}

	return &BzlFile{
		ref:      refs.CFileTable.Insert(c, cfile),
		producer: nil,
		ctx:      c,
	}, nil
}

var _ starlark.Value = &BzlFile{}

type BzlFile struct {
	ref      refs.CFileRef
	producer *Action // TODO(eric): Indirect through a key so we can lazily produce the Action and cache things well
	ctx      *core.Context
}

func (f *BzlFile) String() string {
	val := f.ref.Get(f.ctx).Location.Get(f.ctx)
	return fmt.Sprintf("File(//%s:%s)", val.Pkg.Get(f.ctx), val.Name.Get(f.ctx))
}
func (f *BzlFile) Type() string          { return "file" }
func (f *BzlFile) Freeze()               {}
func (f *BzlFile) Truth() starlark.Bool  { return true }
func (f *BzlFile) Hash() (uint32, error) { return 0, nil }
