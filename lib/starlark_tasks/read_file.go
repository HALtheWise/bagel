package starlark_tasks

import (
	"go.starlark.net/starlark"

	"github.com/HALtheWise/bagel/lib/core"
	"github.com/HALtheWise/bagel/lib/refs"
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

var T_EvalStarlark func(c *core.Context, file refs.LabelRef) StarlarkFileResults

func init() {
	T_EvalStarlark = core.MemoFunc1("T_EvalStarlark",
		func(c *core.Context, file refs.LabelRef) StarlarkFileResults {
			label := refs.LabelTable.Get(c, file)
			thread := &starlark.Thread{Name: "single file thread: " + file.String(), Load: loadFunc(c, label.Pkg)}

			ruleNames := map[string]*BzlRule{}
			thread.SetLocal("rules", ruleNames)
			thread.SetLocal("label", file)

			globals, err := starlark.ExecFile(thread, refs.T_FilepathForLabel(c, file), nil, predeclared)
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

var T_RuleInfoUnconfigured = core.MemoFunc1("T_RuleInfoUnconfigured", func(c *core.Context, l_r refs.LabelRef) *BzlRule {
	l := l_r.Get(c)
	buildfile := refs.T_FindBuildFile(c, l.Pkg)
	if buildfile == core.INVALID {
		return nil
	}

	parsed := T_EvalStarlark(c, buildfile)
	return parsed.rules[l.Name.Get(c)]
})

func loadFunc(c *core.Context, from refs.StringRef) func(*starlark.Thread, string) (starlark.StringDict, error) {
	return func(_ *starlark.Thread, module string) (starlark.StringDict, error) {
		result := T_EvalStarlark(c, refs.ParseRelativeLabel(c, module, from))
		return result.globals, nil
	}
}

func starlarkRuleFunc(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var impl starlark.Callable
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "implementation", &impl); err != nil {
		return nil, err
	}

	rule := &BzlRule{
		DefinedIn: thread.Local("label").(refs.LabelRef),
		Kind:      "",
		Impl:      impl,
		Attrs:     nil,
	}
	return rule, nil
}
