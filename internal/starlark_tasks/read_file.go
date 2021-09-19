package starlark_tasks

import (
	"github.com/HALtheWise/balez/internal/labels"
	"github.com/HALtheWise/balez/internal/task"
	"go.starlark.net/starlark"
)

type StarlarkFileResults struct {
	globals starlark.StringDict
	rules   map[string]*BzlRule
}

var T_EvalStarlark = task.Task1("T_EvalStarlark",
	func(c *task.Context, file labels.Label) StarlarkFileResults {
		thread := &starlark.Thread{Name: "single file thread: " + file.String()}

		ruleNames := map[string]*BzlRule{}
		thread.SetLocal("rules", ruleNames)
		thread.SetLocal("label", file)

		predeclared := starlark.StringDict{
			"rule": starlark.NewBuiltin("rule", starlarkRuleFunc),
		}

		globals, err := starlark.ExecFile(thread, labels.T_FilepathForLabel(c, file), nil, predeclared)
		if err != nil {
			panic(err)
		}

		for _, v := range globals {
			v.Freeze()
		}

		return StarlarkFileResults{globals, ruleNames}
	})

var T_RuleExists = task.Task1("T_RuleExists", func(c *task.Context, l labels.Label) bool {
	buildfile := labels.T_FindBuildFile(c, l.Package)
	if buildfile == labels.NullLabel {
		return false
	}

	parsed := T_EvalStarlark(c, buildfile)
	_, ok := parsed.rules[l.Name]
	return ok
})

func starlarkRuleFunc(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var impl starlark.Callable
	starlark.UnpackArgs(b.Name(), args, kwargs, "_impl", &impl)
	rule := &BzlRule{
		DefinedIn: thread.Local("label").(labels.Label),
		Kind:      "",
		Impl:      impl,
		Attrs:     nil,
	}
	return rule, nil
}
