package analysis

import (
	"fmt"
	"sort"

	"go.starlark.net/starlark"

	"github.com/HALtheWise/bagel/lib/core"
	"github.com/HALtheWise/bagel/lib/loading"
	"github.com/HALtheWise/bagel/lib/refs"
)

type AttrGetterKind int

const (
	ATTR_GETTER AttrGetterKind = iota
	FILE_GETTER
)

// This is ctx.attr or ctx.file
type BzlAttrGetter struct {
	ctx  *BzlCtx
	Kind AttrGetterKind
}

var _ starlark.HasAttrs = &BzlAttrGetter{}

func (a BzlAttrGetter) AttrNames() []string {
	names := make([]string, 0, len(a.ctx.attrs))
	for name, info := range a.ctx.attrs {
		switch a.Kind {
		case ATTR_GETTER:
			names = append(names, name)
		case FILE_GETTER:
			if info.Kind == loading.LABEL_ATTR {
				names = append(names, name)
			}
		default:
			panic("unknown")
		}
	}

	sort.Strings(names)
	return names
}

func (a BzlAttrGetter) Attr(name string) (starlark.Value, error) {
	attr, ok := a.ctx.attrs[name]
	if !ok {
		return nil, nil
	}

	switch attr.Kind {
	case loading.STRING_ATTR:
		if a.Kind == ATTR_GETTER {
			return starlark.String(attr.StringValue), nil
		}
	case loading.LABEL_ATTR:
		return &BzlFile{
			Ref:      attr.Files[0],
			ctx:      a.ctx.ctx,
			producer: nil,
		}, nil
	default:
		panic("Unknown attr kind")
	}
	return nil, nil
}

func (a BzlAttrGetter) String() string        { return "attr getters for " + a.ctx.clabel.String() }
func (a BzlAttrGetter) Type() string          { return "attr" }
func (a BzlAttrGetter) Freeze()               { a.ctx.Freeze() }
func (a BzlAttrGetter) Truth() starlark.Bool  { return starlark.True }
func (a BzlAttrGetter) Hash() (uint32, error) { panic("not implemented") }

type AnalyzedAttrValue struct {
	// Universal
	Kind loading.AttrKind

	// Kind-specific
	StringValue string
	Files       []refs.CFileRef
}

func getAttrs(c *core.Context, infos []loading.AttrInfo, values []loading.AttrValue) map[string]AnalyzedAttrValue {

	if len(infos) != len(values) {
		panic("Attr length mismatch")
	}
	result := make(map[string]AnalyzedAttrValue, len(infos))
	for i, info := range infos {
		value := AnalyzedAttrValue{
			Kind: info.Kind,
		}

		switch info.Kind {
		case loading.STRING_ATTR:
			value.StringValue = values[i].StringValue
		case loading.LABEL_ATTR:
			// Recursively compute file info
			label := values[i].LabelValue

			// First, check if this names a rule
			cLabel := refs.CLabelTable.Insert(c, refs.CLabel{
				Label: label,
				// TODO: dynamic config
				Config: refs.TARGET_CONFIG})

			target := T_AnalyzeTarget(c, cLabel)

			if target != nil {
				// panic("We don't yet support depending on other targets :(")

				files, err := GetDefaultFiles(target.Providers)
				if err != nil {
					panic(err)
				}
				value.Files = files
			} else {
				// This must be an on-disk file
				value.Files = []refs.CFileRef{
					refs.CFileTable.Insert(c,
						refs.CFile{
							Location: label,
							Source:   refs.FILESYSTEM_SOURCE,
						})}
			}
			fmt.Println("Got label", value.Files)

		default:
			panic("Unknown kind")
		}

		result[info.Name] = value
	}
	return result
}
