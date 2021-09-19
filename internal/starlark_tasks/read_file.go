package starlark_tasks

import (
	"fmt"

	"github.com/HALtheWise/balez/internal/task"
	"go.starlark.net/starlark"
)

type StarlarkFileResults struct {
	globals starlark.StringDict
	rules   map[string]*BzlRule
}

var T_ExecuteFile = task.Task1("starlark.ExecuteFileT",
	func(c *task.Context, path string) StarlarkFileResults {
		thread := &starlark.Thread{Name: "single file thread: " + path}

		ruleNames := map[string]*BzlRule{}
		thread.SetLocal("rules", ruleNames)
		thread.SetLocal("label", Label{"", path, ""})

		predeclared := starlark.StringDict{
			"rule": starlark.NewBuiltin("rule", starlarkRuleFunc),
		}

		globals, err := starlark.ExecFile(thread, path, nil, predeclared)
		if err != nil {
			panic(err)
		}

		for _, v := range globals {
			v.Freeze()
		}

		fmt.Printf("Globals=%v, rules=%v", globals, ruleNames)

		return StarlarkFileResults{globals, ruleNames}
	})

func starlarkRuleFunc(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var impl starlark.Callable
	starlark.UnpackArgs(b.Name(), args, kwargs, "_impl", &impl)
	rule := &BzlRule{
		DefinedIn: thread.Local("label").(Label),
		Kind:      "",
		Impl:      impl,
		Attrs:     nil,
	}
	return rule, nil
}
