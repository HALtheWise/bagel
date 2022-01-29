package loading

import (
	"go.starlark.net/starlark"
)

type AttrKind int

const STRING_ATTR AttrKind = iota

type AttrInfo struct {
	OpaqueObject
	// Universal parameters
	Kind      AttrKind
	Name      string
	Doc       string
	Mandatory bool

	// Kind-specific (maybe replace with interface?)
	StringDefault string
}

// AttrValue represents a specific value provided to an attribute at the end of the loading stage.
// Eventually, this will need to handle select() as well.
type AttrValue struct {
	// Universal
	Kind AttrKind

	// Kind-specific
	StringValue string
}

// https://docs.bazel.build/versions/main/skylark/lib/attr.html
var BzlAttrs = NewBuiltinStruct("attrs", BuiltinStructMembers{
	"string": stringAttr,
})

func stringAttr(thread *starlark.Thread, fn *starlark.Builtin,
	args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {

	mandatory := false
	Default := ""
	doc := ""
	starlark.UnpackArgs("string", args, kwargs, "default?", &Default, "doc?", &doc, "mandatory?", &mandatory)

	return &AttrInfo{
		Kind:          STRING_ATTR,
		Mandatory:     mandatory,
		Doc:           doc,
		StringDefault: Default,
	}, nil
}
