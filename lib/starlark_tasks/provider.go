package starlark_tasks

import (
	"fmt"

	"go.starlark.net/starlark"
)

func NewBuiltinProvider(name string) *MetaProvider {
	return &MetaProvider{name: name}
}

func starlarkProviderFunc(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Callable, error) {
	if err := starlark.UnpackArgs("provider", args, kwargs); err != nil {
		return nil, err
	}

	return &MetaProvider{}, nil
}

// MetaProvider represents the result of a provider() call (or a builtin like DefaultInfo)
type MetaProvider struct {
	name string // Populated when the module's done evaluating
}

func (p *MetaProvider) String() string        { return p.name }
func (p *MetaProvider) Type() string          { return "provider" }
func (p *MetaProvider) Freeze()               {}
func (p *MetaProvider) Truth() starlark.Bool  { return true }
func (p *MetaProvider) Hash() (uint32, error) { return starlark.String(p.name).Hash() }
func (p *MetaProvider) Name() string          { return p.name }

func (p *MetaProvider) Calllib(thread *starlark.Thread, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	if len(args) != 0 {
		return nil, fmt.Errorf("%s does not take positional args", p.Name())
	}

	result := &Provider{kind: p, Data: make(starlark.StringDict)}

	for _, kwarg := range kwargs {
		if kwarg.Len() != 2 {
			panic("wtf")
		}
		result.Data[string(kwarg.Index(0).(starlark.String))] = kwarg.Index(1)
	}

	return result, nil
}

type Provider struct {
	kind *MetaProvider
	Data starlark.StringDict
}

func (p *Provider) String() string { return fmt.Sprintf("%s%s", p.kind.name, p.Data) }
func (p *Provider) Type() string   { return "depset" }
func (p *Provider) Freeze() {
	for _, item := range p.Data {
		item.Freeze()
	}
}
func (p *Provider) Truth() starlark.Bool  { return true }
func (p *Provider) Hash() (uint32, error) { panic("not implemented") }
