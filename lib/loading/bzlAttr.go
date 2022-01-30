package loading

import (
	"errors"

	"go.starlark.net/starlark"

	"github.com/HALtheWise/bagel/lib/refs"
)

type AttrKind int

const (
	STRING_ATTR AttrKind = iota
	// STRING_DICT_ATTR
	// STRING_LIST_ATTR
	// STRING_LIST_DICT_ATTR
	// BOOL_ATTR
	// INT_ATTR
	// INT_LIST_ATTR
	LABEL_ATTR
	// LABEL_KEYED_STRING_DICT_ATTR
	// LABEL_LIST_ATTR
	// OUTPUT_ATTR
	// OUTPUT_LIST_ATTR
)

type AttrInfo struct {
	OpaqueObject
	// Universal parameters
	Kind      AttrKind
	Name      string
	Doc       string
	Mandatory bool

	// Kind-specific defaults (maybe use any?)
	StringDefault string

	// Info for Labels
	AllowFiles     []string // If nil, do not allow files. [""] allows any file.
	SingleFileOnly bool     // Don't allow multi-file inputs
	// Cfg             refs.ConfigRef
}

// AttrValue represents a specific value provided to an attribute at the end of the loading stage.
// Eventually, this will need to handle select() as well.
type AttrValue struct {
	// Universal
	Kind AttrKind

	// Kind-specific
	StringValue string // TODO: StringRef?
	LabelValue  refs.LabelRef
}

// https://docs.bazel.build/versions/main/skylark/lib/attr.html
var BzlAttrs = NewBuiltinStruct("attrs", BuiltinStructMembers{
	"string": attrGetter(STRING_ATTR),
	"label":  attrGetter(LABEL_ATTR),
})

type stringSlice []string

func (s *stringSlice) Unpack(v starlark.Value) error {
	switch v := v.(type) {
	case starlark.Bool:
		if v == starlark.True {
			*s = []string{""}
		} else {
			*s = []string{}
		}
	case *starlark.List:
		*s = make(stringSlice, 0, v.Len())
		for i := 0; i < v.Len(); i++ {
			val, ok := v.Index(i).(starlark.String)
			if !ok {
				return errors.New("Expected strings")
			}
			*s = append(*s, string(val))
		}
	}
	return nil
}

func attrGetter(kind AttrKind) func(*starlark.Thread, *starlark.Builtin,
	starlark.Tuple, []starlark.Tuple) (starlark.Value, error) {

	return func(thread *starlark.Thread, fn *starlark.Builtin,
		args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {

		if len(args) > 0 {
			return nil, errors.New("Positional args not supported for attrs")
		}

		result := AttrInfo{Kind: kind}

		// Universal
		unpackPairs := []any{"doc?", &result.Doc, "mandatory?", &result.Mandatory}

		// Default
		switch kind {
		case STRING_ATTR, LABEL_ATTR:
			unpackPairs = append(unpackPairs, "default?", &result.StringDefault)
		}

		// Label info
		var allow_files, allow_single_file stringSlice
		switch kind {
		case LABEL_ATTR:
			unpackPairs = append(unpackPairs,
				// "executable", &result.x,
				"allow_files", &allow_files,
				"allow_single_file?", &allow_single_file,
				// "providers", &result.x,
				// "cfg?", &result.Cfg,
				// "aspects", &result.x,
			)
		}

		starlark.UnpackArgs("string", args, kwargs, unpackPairs...)

		if allow_files != nil {
			if allow_single_file != nil {
				return nil, errors.New("incompatable")
			}
			result.AllowFiles = []string(allow_files)
		} else if allow_single_file != nil {
			result.AllowFiles = []string(allow_single_file)
			result.SingleFileOnly = true
		}

		return &result, nil
	}
}
