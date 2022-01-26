package analysis

import (
	"fmt"

	"go.starlark.net/starlark"
)

const kActionsKey = "__actions__"

type BzlActions struct {
	ctx *BzlCtx
}

func (a *BzlActions) AttrNames() []string {
	return []string{"write", "declare_file"}
	// TODO: "args", "declare_directory", "declare_symlink", "do_nothing", "expand_template", "run", "run_shell", "symlink"
}

func (a *BzlActions) Attr(name string) (starlark.Value, error) {
	switch name {
	case "write":
		return starlark.NewBuiltin("write", actionWrite), nil
	case "declare_file":
		builtin := starlark.NewBuiltin("declare_file", actionDeclareFile)
		builtin = builtin.BindReceiver(a)
		return builtin, nil
	}
	return nil, nil
}

func (a *BzlActions) String() string        { return "actions for " + a.ctx.label.String() }
func (a *BzlActions) Type() string          { return "actions" }
func (a *BzlActions) Freeze()               {}
func (a *BzlActions) Truth() starlark.Bool  { return starlark.True }
func (a *BzlActions) Hash() (uint32, error) { panic("not implemented") }

func actionWrite(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var output *File
	var content string
	if err := starlark.UnpackArgs("write", args, kwargs, "output", &output, "content", &content); err != nil {
		return nil, err
	}

	write := &Action{
		kind:         WRITE,
		inputs:       nil,
		outputs:      []*File{output},
		writeContent: content,
	}

	if err := registerAction(thread, write); err != nil {
		return nil, err
	}

	return starlark.None, nil
}

func registerAction(thread *starlark.Thread, action *Action) error {
	for _, output := range action.outputs {
		if output.producer != nil {
			return fmt.Errorf("File cannot be output of multiple actions")
		}
		output.producer = action
	}

	actions, _ := thread.Local(kActionsKey).([]*Action)
	thread.SetLocal(kActionsKey, append(actions, action))

	return nil
}

type ActionKind int8

const (
	WRITE ActionKind = iota
)

type Action struct {
	kind            ActionKind
	inputs, outputs []*File

	writeContent string
}
