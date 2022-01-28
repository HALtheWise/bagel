package loading

import (
	"go.starlark.net/starlark"

	"github.com/HALtheWise/bagel/lib/core"
	"github.com/HALtheWise/bagel/lib/refs"
)

type Target struct {
	Rule   *BzlRule
	Args   starlark.Tuple
	Kwargs []starlark.Tuple
}

type StarlarkFileResults struct {
	globals starlark.StringDict
	targets map[string]*Target
}

var predeclared = starlark.StringDict{
	"rule":        starlark.NewBuiltin("rule", starlarkRuleFunc),
	"depset":      starlark.NewBuiltin("depset", starlarkDepsetFunc),
	"DefaultInfo": NewBuiltinProvider("DefaultInfo"),
}

var T_EvalStarlark func(c *core.Context, file refs.LabelRef) StarlarkFileResults

func init() {
	T_EvalStarlark = core.Task1("T_EvalStarlark",
		func(c *core.Context, file refs.LabelRef) StarlarkFileResults {
			label := refs.LabelTable.Get(c, file)
			thread := &starlark.Thread{Name: "single file thread: " + file.String(), Load: LoadFunc(c, label.Pkg)}

			targets := map[string]*Target{}
			thread.SetLocal(kTargetsKey, targets)
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

			return StarlarkFileResults{globals, targets}
		})
}

var T_LoadTarget = core.Task1("T_LoadTarget", func(c *core.Context, l_r refs.LabelRef) *Target {
	l := l_r.Get(c)
	buildfile := refs.T_FindBuildFile(c, l.Pkg)
	if buildfile == core.INVALID {
		return nil
	}

	parsed := T_EvalStarlark(c, buildfile)
	return parsed.targets[l.Name.Get(c)]
})

func LoadFunc(c *core.Context, from refs.PackageRef) func(*starlark.Thread, string) (starlark.StringDict, error) {
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
