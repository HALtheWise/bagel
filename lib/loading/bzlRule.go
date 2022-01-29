package loading

import (
	"fmt"

	"go.starlark.net/starlark"

	"github.com/HALtheWise/bagel/lib/refs"
)

const kTargetsKey = "__targets__"

var _ starlark.Callable = &BzlRule{}

type BzlRule struct {
	DefinedIn refs.LabelRef
	Kind      string
	Impl      starlark.Callable
	Attrs     []Attr
}

func (r *BzlRule) String() string {
	if r.Kind != "" {
		return r.Kind
	} else {
		return fmt.Sprintf("%s defined in %s", r.Name(), &r.DefinedIn)
	}
}

func (r *BzlRule) Type() string          { return "builtin" }
func (r *BzlRule) Truth() starlark.Bool  { return true }
func (r *BzlRule) Hash() (uint32, error) { return 0, fmt.Errorf("rule objects are not hashable") }

func (r *BzlRule) Freeze() {
	if r.Impl != nil {
		r.Impl.Freeze()
	}
}

// Invoked when you call rule(impl=...) in Starlark
func starlarkRuleFunc(thread *starlark.Thread, b *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var impl starlark.Callable
	var attrs *starlark.Dict
	if err := starlark.UnpackArgs(b.Name(), args, kwargs, "implementation", &impl, "attrs?", &attrs); err != nil {
		return nil, err
	}

	rule := &BzlRule{
		DefinedIn: thread.Local("label").(refs.LabelRef),
		Kind:      "",
		Impl:      impl,
		Attrs:     nil,
	}

	// parse attrs={...}
	for _, k := range attrs.Keys() {
		key := k.(starlark.String)
		v, ok, err := attrs.Get(key)
		if !ok || err != nil {
			return nil, fmt.Errorf("Unable to get %v: %w", key, err)
		}
		switch val := v.(type) {
		case *Attr:
			copy := *val
			copy.Name = string(key)
			rule.Attrs = append(rule.Attrs, copy)
		default:
			return nil, fmt.Errorf("Unknown attr %s (%T)", v, v)
		}
	}

	return rule, nil
}

func (r *BzlRule) Name() string {
	if r.Kind != "" {
		return r.Kind
	} else {
		return "anonymous rule"
	}
}

func (r *BzlRule) CallInternal(thread *starlark.Thread, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var name string

	parseArgs := []any{"name", &name}
	attrs := make([]starlark.Value, len(r.Attrs))
	for i, attr := range r.Attrs {
		if attr.Mandatory {
			parseArgs = append(parseArgs, attr.Name)
		} else {
			parseArgs = append(parseArgs, attr.Name+"?")
		}
		parseArgs = append(parseArgs, &attrs[i])
	}

	if err := starlark.UnpackArgs(r.Name(), args, kwargs, parseArgs...); err != nil {
		return nil, err
	}

	targets := thread.Local(kTargetsKey).(map[string]*Target)
	targets[name] = &Target{r, attrs}
	return starlark.None, nil
}
