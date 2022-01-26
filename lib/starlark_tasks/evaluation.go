package starlark_tasks

import (
	"fmt"

	"go.starlark.net/starlark"

	"github.com/HALtheWise/bagel/lib/core"
	"github.com/HALtheWise/bagel/lib/refs"
)

type EvaluatedRule struct {
	Name, Kind refs.StringRef
	Providers  []*Provider
}

func (c *EvaluatedRule) String() string {
	return fmt.Sprintf("%s{\"%s\", provides:%s}", c.Kind, c.Name, c.Providers)
}

var T_RuleInfoEvaluated = core.MemoFunc1("T_RuleInfoEvaluated", func(c *core.Context, label refs.LabelRef) *EvaluatedRule {
	label_v := label.Get(c)

	unconfigured := T_RuleInfoUnconfigured(c, label)
	if unconfigured == nil {
		return nil
	}

	thread := &starlark.Thread{Name: "Rule evaluation thread: " + label.String(), Load: loadFunc(c, label_v.Pkg)}
	bzlResult, err := starlark.Call(thread, unconfigured.Impl, starlark.Tuple{
		&BzlCtx{BzlLabel{label: label, frozen: false, ctx: c}, c}, // ctx
	}, []starlark.Tuple{})
	if err != nil {
		panic(err)
	}

	var providers []*Provider

	if seq, ok := bzlResult.(starlark.Indexable); ok {
		for i := 0; i < seq.Len(); i++ {
			providers = append(providers, seq.Index(i).(*Provider))
		}
	}

	return &EvaluatedRule{
		Kind:      refs.StringTable.Insert(c, unconfigured.Kind),
		Name:      label_v.Name,
		Providers: providers,
	}
})
