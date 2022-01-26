package loading

import (
	"fmt"

	"go.starlark.net/starlark"

	"github.com/HALtheWise/bagel/lib/refs"
)

var _ starlark.Callable = &BzlRule{}

type Attr struct {
	kind string
}

type BzlRule struct {
	DefinedIn refs.LabelRef
	Kind      string
	Impl      starlark.Callable
	Attrs     map[string]Attr
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

func (r *BzlRule) Name() string {
	if r.Kind != "" {
		return r.Kind
	} else {
		return "anonymous rule"
	}
}

func (r *BzlRule) CallInternal(thread *starlark.Thread, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var name string

	if err := starlark.UnpackArgs(r.Name(), args, kwargs, "name", &name); err != nil {
		return nil, err
	}

	rules := thread.Local("rules").(map[string]*BzlRule)
	rules[name] = r
	return starlark.None, nil
}
