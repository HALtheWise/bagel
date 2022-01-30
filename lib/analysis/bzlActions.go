package analysis

import (
	"fmt"

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
	case "expand_template":
		builtin := starlark.NewBuiltin("expand_template", actionExpandTemplate)
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
		Kind:         WRITE,
		Inputs:       nil,
		Outputs:      []refs.CFileRef{output.Ref},
		Executable:   is_executable,
		WriteContent: content,
	}

	actions.ctx.actions = append(actions.ctx.actions, write)

	return starlark.None, nil
}

func actionExpandTemplate(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var template *BzlFile
	var output *BzlFile
	var rawSubs *starlark.Dict
	var is_executable bool = false
	if err := starlark.UnpackArgs("expand_template", args, kwargs,
		"template", &template,
		"output", &output,
		"substitutions", &rawSubs,
		"is_executable?", &is_executable); err != nil {
		return nil, err
	}

	actions := fn.Receiver().(BzlActions)

	parsedSubs := make(map[string]string)
	for _, kv := range rawSubs.Items() {
		key, ok1 := kv.Index(0).(starlark.String)
		value, ok2 := kv.Index(1).(starlark.String)
		if !ok1 || !ok2 {
			return nil, fmt.Errorf("Invalid dict item: %s", kv)
		}
		parsedSubs[string(key)] = string(value)
	}

	write := &Action{
		Kind:                        EXPAND_TEMPLATE,
		Inputs:                      []refs.CFileRef{template.Ref},
		Outputs:                     []refs.CFileRef{output.Ref},
		Executable:                  is_executable,
		ExpandTemplateSubstitutions: parsedSubs,
	}

	actions.ctx.actions = append(actions.ctx.actions, write)

	return starlark.None, nil
}
