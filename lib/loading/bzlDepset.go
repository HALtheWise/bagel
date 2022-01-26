package loading

import (
	"fmt"

	"go.starlark.net/starlark"
)

func starlarkDepsetFunc(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var direct *starlark.List
	if err := starlark.UnpackArgs("depset", args, kwargs, "direct", &direct); err != nil {
		return nil, err
	}

	result := &Depset{}
	for i := 0; i < direct.Len(); i++ {
		result.Items = append(result.Items, direct.Index(i))
	}

	return result, nil
}

type Depset struct {
	// TODO: support efficient merging
	Items []starlark.Value
}

func (d *Depset) String() string { return fmt.Sprint(d.Items) }
func (d *Depset) Type() string   { return "depset" }
func (d *Depset) Freeze() {
	for _, item := range d.Items {
		item.Freeze()
	}
}
func (d *Depset) Truth() starlark.Bool  { return len(d.Items) > 0 }
func (d *Depset) Hash() (uint32, error) { panic("Cannot hash depset") }
