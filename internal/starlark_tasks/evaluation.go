package starlark_tasks

import (
	"fmt"

	"github.com/HALtheWise/bales/internal/labels"
	"github.com/HALtheWise/bales/internal/task"
	"go.starlark.net/starlark"
)

type EvaluatedRule struct {
	Name, Kind string
	Providers  []*Provider
}

func (c *EvaluatedRule) String() string {
	return fmt.Sprintf("%s{\"%s\", provides:%s}", c.Kind, c.Name, c.Providers)
}

var T_RuleInfoEvaluated = task.Task1("T_RuleInfoEvaluated", func(c *task.Context, label labels.Label) *EvaluatedRule {
	unconfigured := T_RuleInfoUnconfigured(c, label)
	if unconfigured == nil {
		return nil
	}

	thread := &starlark.Thread{Name: "Rule evaluation thread: " + label.String(), Load: loadFunc(c, label.Package)}
	bzlResult, err := starlark.Call(thread, unconfigured.Impl, starlark.Tuple{
		&BzlCtx{BzlLabel{label: label, frozen: false}}, // ctx
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
		Kind:      unconfigured.Kind,
		Name:      label.Name,
		Providers: providers,
	}
})
