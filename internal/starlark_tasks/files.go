package starlark_tasks

import (
	"github.com/HALtheWise/balez/internal/labels"
	"go.starlark.net/starlark"
)

func actionDeclareFile(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var filename string
	if err := starlark.UnpackArgs("declare_file", args, kwargs, "filename", &filename); err != nil {
		return nil, err
	}

	actions := fn.Receiver().(*BzlActions)
	label := labels.Label{Package: actions.ctx.label.label.Package, Name: filename}

	return &File{path: label, producer: nil}, nil
}

type File struct {
	path     labels.Label
	producer *Action // TODO(eric): Indirect through a key so we can lazily produce the Action and cache things well
}

func (f *File) String() string        { return f.path.Name }
func (f *File) Type() string          { return "file" }
func (f *File) Freeze()               {}
func (f *File) Truth() starlark.Bool  { return true }
func (f *File) Hash() (uint32, error) { return 0, nil }
