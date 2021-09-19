package starlark_tasks

import (
	"fmt"

	"github.com/HALtheWise/balez/internal/task"
	"go.starlark.net/starlark"
)

type StarlarkFileResults struct {
	globals starlark.StringDict
	rules   map[string]bool
}

var ExecuteFileT = task.Task1("starlark.ExecuteFileT",
	func(c *task.Context, path string) StarlarkFileResults {
		thread := &starlark.Thread{Name: "single file thread: " + path}

		ruleNames := []string{}
		thread.SetLocal("rules", &ruleNames)

		predeclared := starlark.StringDict{
			"rule": starlark.NewBuiltin("rule", starlarkRuleFunc),
		}

		globals, err := starlark.ExecFile(thread, path, nil, predeclared)

		fmt.Printf("Globals, err, %v, %v", globals, err)

		return StarlarkFileResults{}
	})

func starlarkRuleFunc(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {

	rule := func(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
		var name string
		starlark.UnpackArgs(b.Name(), args, kwargs, "name", &name)

		rules := thread.Local("rules").(*[]string)
		*rules = append(*rules, name)
		return nil, nil
	}

	return starlark.NewBuiltin("unnamed", rule), nil
}
