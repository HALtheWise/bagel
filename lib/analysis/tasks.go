package analysis

import (
	"go.starlark.net/starlark"

	"github.com/HALtheWise/bagel/lib/core"
	"github.com/HALtheWise/bagel/lib/loading"
	"github.com/HALtheWise/bagel/lib/refs"
)

type AnalyzedTarget struct {
	Name, Kind refs.StringRef
	Providers  []*loading.Provider
	Actions    []*Action
}

var T_AnalyzeTarget = core.Task1("T_AnalyzeTarget", func(c *core.Context, label refs.LabelRef) *AnalyzedTarget {
	label_v := label.Get(c)

	unconfigured := loading.T_LoadTarget(c, label)
	if unconfigured == nil {
		return nil
	}

	thread := &starlark.Thread{Name: "Rule evaluation thread: " + label.String(), Load: loading.LoadFunc(c, label_v.Pkg)}
	bzlCtx := &BzlCtx{ctx: c, label: label}
	bzlResult, err := starlark.Call(thread, unconfigured.Rule.Impl, starlark.Tuple{bzlCtx}, nil)
	if err != nil {
		panic(err)
	}

	var providers []*loading.Provider

	if seq, ok := bzlResult.(starlark.Indexable); ok {
		for i := 0; i < seq.Len(); i++ {
			providers = append(providers, seq.Index(i).(*loading.Provider))
		}
	}

	actions := bzlCtx.actions

	return &AnalyzedTarget{
		Kind:      refs.StringTable.Insert(c, unconfigured.Rule.Kind),
		Name:      label_v.Name,
		Providers: providers,
		Actions:   actions,
	}
})
