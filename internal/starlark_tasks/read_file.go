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

var predeclared = starlark.StringDict{
	"rule":        starlark.NewBuiltin("rule", starlarkRuleFunc),
	"depset":      starlark.NewBuiltin("depset", starlarkDepsetFunc),
	"DefaultInfo": NewBuiltinProvider("DefaultInfo"),
}

var T_EvalStarlark func(c *task.Context, file labels.Label) StarlarkFileResults

func init() {
	T_EvalStarlark = task.Task1("T_EvalStarlark",
		func(c *task.Context, file labels.Label) StarlarkFileResults {
			thread := &starlark.Thread{Name: "single file thread: " + file.String(), Load: loadFunc(c, file.Package)}

			ruleNames := map[string]*BzlRule{}
			thread.SetLocal("rules", ruleNames)
			thread.SetLocal("label", file)

			globals, err := starlark.ExecFile(thread, labels.T_FilepathForLabel(c, file), nil, predeclared)
			if err != nil {
				panic(err)
			}

			// Store the names of the rules. This only works if they were uniquely assigned.
			for name, v := range globals {
				if rule, ok := v.(*BzlRule); ok {
					rule.Kind = name
				}
			}

			// Freeze everything before we throw out the thread
			for _, v := range globals {
				v.Freeze()
			}

			return StarlarkFileResults{globals, ruleNames}
		})
}

var T_RuleInfoUnconfigured = task.Task1("T_RuleInfoUnconfigured", func(c *task.Context, l labels.Label) *BzlRule {
	buildfile := labels.T_FindBuildFile(c, l.Package)
	if buildfile == labels.NullLabel {
		return nil
	}

	parsed := T_EvalStarlark(c, buildfile)
	return parsed.rules[l.Name]
})

func loadFunc(c *task.Context, from labels.Package) func(*starlark.Thread, string) (starlark.StringDict, error) {
	return func(_ *starlark.Thread, module string) (starlark.StringDict, error) {
		result := T_EvalStarlark(c, labels.ParseRelativeLabel(module, from))
		return result.globals, nil
	}
}

func starlarkRuleFunc(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var impl starlark.Callable
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "implementation", &impl); err != nil {
		return nil, err
	}

	rule := &BzlRule{
		DefinedIn: thread.Local("label").(labels.Label),
		Kind:      "",
		Impl:      impl,
		Attrs:     nil,
	}
	return rule, nil
}
