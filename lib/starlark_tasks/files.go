package starlark_tasks

import (
	"fmt"

	"go.starlark.net/starlark"

	"github.com/HALtheWise/bagel/lib/core"
	"github.com/HALtheWise/bagel/lib/refs"
)

func actionDeclareFile(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var filename string
	if err := starlark.UnpackArgs("declare_file", args, kwargs, "filename", &filename); err != nil {
		return nil, err
	}

	actions := fn.Receiver().(*BzlActions)
	c := actions.ctx.ctx
	label := refs.Label{Pkg: actions.ctx.label.label.Get(c).Pkg, Name: refs.StringTable.Insert(c, filename)}

	return &File{
		path:     refs.LabelTable.Insert(c, label),
		producer: nil}, nil
}

type File struct {
	path     refs.LabelRef
	producer *Action // TODO(eric): Indirect through a key so we can lazily produce the Action and cache things well
	ctx      *core.Context
}

func (f *File) String() string {
	val := f.path.Get(f.ctx)
	return fmt.Sprintf("File(//{}:{})", val.Pkg.Get(f.ctx), val.Name.Get(f.ctx))
}
func (f *File) Type() string          { return "file" }
func (f *File) Freeze()               {}
func (f *File) Truth() starlark.Bool  { return true }
func (f *File) Hash() (uint32, error) { return 0, nil }
