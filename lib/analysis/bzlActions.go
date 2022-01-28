package analysis

import (
	"go.starlark.net/starlark"

	"github.com/HALtheWise/bagel/lib/refs"
)

// https://docs.bazel.build/versions/main/skylark/lib/actions.html
var _ starlark.HasAttrs = BzlActions{}

type BzlActions struct {
	ctx *BzlCtx
}

func (a BzlActions) AttrNames() []string {
	return []string{"write", "declare_file"}
	// TODO: "args", "declare_directory", "declare_symlink", "do_nothing", "expand_template", "run", "run_shell", "symlink"
}

func (a BzlActions) Attr(name string) (starlark.Value, error) {
	switch name {
	case "write":
		builtin := starlark.NewBuiltin("write", actionWrite)
		builtin = builtin.BindReceiver(a)
		return builtin, nil
	case "declare_file":
		builtin := starlark.NewBuiltin("declare_file", actionDeclareFile)
		builtin = builtin.BindReceiver(a)
		return builtin, nil
	}
	return nil, nil
}

func (a BzlActions) String() string        { return "actions for " + a.ctx.clabel.String() }
func (a BzlActions) Type() string          { return "actions" }
func (a BzlActions) Freeze()               { a.ctx.Freeze() }
func (a BzlActions) Truth() starlark.Bool  { return starlark.True }
func (a BzlActions) Hash() (uint32, error) { panic("not implemented") }

func actionWrite(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var output *BzlFile
	var content string
	var is_executable bool
	if err := starlark.UnpackArgs("write", args, kwargs,
		"output", &output,
		"content", &content,
		"is_executable?", &is_executable); err != nil {
		return nil, err
	}

	actions := fn.Receiver().(BzlActions)

	write := &Action{
		kind:          WRITE,
		inputs:        nil,
		outputs:       []refs.CFileRef{output.ref},
		is_executable: is_executable,
		writeContent:  content,
	}

	actions.ctx.actions = append(actions.ctx.actions, write)

	return starlark.None, nil
}
