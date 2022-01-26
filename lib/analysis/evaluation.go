package analysis

import (
	"fmt"

	"go.starlark.net/starlark"

	"github.com/HALtheWise/bagel/lib/core"
	"github.com/HALtheWise/bagel/lib/loading"
	"github.com/HALtheWise/bagel/lib/refs"
)

type AnalyzedRule struct {
	Name, Kind refs.StringRef
	Providers  []*loading.Provider
}

func (c *AnalyzedRule) String() string {
	return fmt.Sprintf("%s{\"%s\", provides:%s}", c.Kind, c.Name, c.Providers)
}

var T_AnalyzeRule = core.Task1("T_AnalyzeRule", func(c *core.Context, label refs.LabelRef) *AnalyzedRule {
	label_v := label.Get(c)

	unconfigured := loading.T_RuleInfoUnconfigured(c, label)
	if unconfigured == nil {
		return nil
	}

	thread := &starlark.Thread{Name: "Rule evaluation thread: " + label.String(), Load: loading.LoadFunc(c, label_v.Pkg)}
	bzlResult, err := starlark.Call(thread, unconfigured.Impl, starlark.Tuple{
		&BzlCtx{BzlLabel{label: label, frozen: false, ctx: c}, c}, // ctx
	}, []starlark.Tuple{})
	if err != nil {
		panic(err)
	}

	var providers []*loading.Provider

	if seq, ok := bzlResult.(starlark.Indexable); ok {
		for i := 0; i < seq.Len(); i++ {
			providers = append(providers, seq.Index(i).(*loading.Provider))
		}
	}

	return &AnalyzedRule{
		Kind:      refs.StringTable.Insert(c, unconfigured.Kind),
		Name:      label_v.Name,
		Providers: providers,
	}
})
