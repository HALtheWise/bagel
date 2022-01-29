package loading

import (
	"errors"
	"sort"

	"go.starlark.net/starlark"
)

// A BzlBuiltinStruct holds a set of function pointers that are namespaced together.
type BzlBuiltinStruct struct {
	name    string
	members map[string]*starlark.Builtin
}

type BuiltinStructMembers map[string]func(thread *starlark.Thread, fn *starlark.Builtin,
	args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error)

func NewBuiltinStruct(name string, members BuiltinStructMembers) *BzlBuiltinStruct {
	result := BzlBuiltinStruct{name: name, members: make(map[string]*starlark.Builtin)}
	for k, v := range members {
		result.members[k] = starlark.NewBuiltin(k, v)
	}
	return &result
}

func (b *BzlBuiltinStruct) AttrNames() []string {
	names := make([]string, 0, len(b.members))
	for n := range b.members {
		names = append(names, n)
	}
	sort.Strings(names)
	return names
}

func (b *BzlBuiltinStruct) Attr(name string) (starlark.Value, error) {
	if m, ok := b.members[name]; ok {
		return m, nil
	}
	return nil, nil
}

func (b *BzlBuiltinStruct) String() string       { return b.name }
func (b *BzlBuiltinStruct) Type() string         { return "builtin_struct" }
func (b *BzlBuiltinStruct) Freeze()              {}
func (b *BzlBuiltinStruct) Truth() starlark.Bool { return true }
func (b *BzlBuiltinStruct) Hash() (uint32, error) {
	return 0, errors.New("builtin_struct is unhashable")
}

// An OpaqueObject represents a nonserializable unfreezable builtin
type OpaqueObject struct{}

var _ starlark.Value = &OpaqueObject{}

func (o *OpaqueObject) String() string        { return "<builtin_object>" }
func (o *OpaqueObject) Type() string          { return "<builtin_object>" }
func (o *OpaqueObject) Freeze()               { panic("Unable to freeze builtin object") }
func (o *OpaqueObject) Truth() starlark.Bool  { return true }
func (o *OpaqueObject) Hash() (uint32, error) { return 0, errors.New("builtin_object is unhashable") }
