package loading

import (
	"go.starlark.net/starlark"

	"github.com/HALtheWise/bagel/lib/core"
	"github.com/HALtheWise/bagel/lib/refs"
)

// Targets are the primary output of the Loading stage into the Analysis stage.
// Eventually we should figure out how to make these serializable.
type Target struct {
	Rule  *BzlRule
	Attrs map[string]AttrValue
}

type StarlarkFileResults struct {
	globals starlark.StringDict
	targets map[string]*Target
}

var DefaultInfo = NewBuiltinProvider("DefaultInfo")

var predeclared = starlark.StringDict{
	"rule":        starlark.NewBuiltin("rule", starlarkRuleFunc),
	"depset":      starlark.NewBuiltin("depset", starlarkDepsetFunc),
	"attr":        BzlAttrs,
	"DefaultInfo": DefaultInfo,
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
	if buildfile == refs.INVALID_LABEL {
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
