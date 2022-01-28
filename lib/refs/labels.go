package refs

import (
	"fmt"

	"github.com/HALtheWise/bagel/lib/core"
)

// StringRef is used to refer to generic "strings" in various forms
type StringRef uint32

const (
	EMPTYSTRING StringRef = iota
	INVALIDSTRING
)

var StringTable = core.NewInternTable(map[string]StringRef{
	"":                      EMPTYSTRING,
	"__invalid\x00string__": INVALIDSTRING,
})

func (r StringRef) Get(c *core.Context) string { return StringTable.Get(c, r) }

type PackageRef uint32

const (
	INVALID_PACKAGE PackageRef = iota
)

type Package struct {
	Workspace StringRef
	RelPath   StringRef
}

var PackageTable = core.NewInternTable(map[Package]PackageRef{
	{INVALIDSTRING, INVALIDSTRING}: INVALID_PACKAGE,
})

func (r PackageRef) Get(c *core.Context) Package { return PackageTable.Get(c, r) }

type LabelRef uint32

const (
	INVALID_LABEL LabelRef = iota
)

type Label struct {
	Pkg  PackageRef
	Name StringRef
}

var LabelTable = core.NewInternTable(map[Label]LabelRef{
	{INVALID_PACKAGE, INVALIDSTRING}: INVALID_LABEL,
})

func (r LabelRef) Get(c *core.Context) Label { return LabelTable.Get(c, r) }

// String functions

func (r StringRef) String() string { return fmt.Sprintf(`r"%s"`, r.Get(core.DefaultContext)) }

func (r PackageRef) String() string {
	c := core.DefaultContext
	val := r.Get(c)
	return fmt.Sprintf("r%+v", val)
}
func (p Package) String() string {
	c := core.DefaultContext
	return fmt.Sprintf("@%s//%s", p.Workspace.Get(c), p.RelPath.Get(c))
}

func (r LabelRef) String() string { return fmt.Sprintf("r%+v", r.Get(core.DefaultContext)) }

func (l Label) String() string {
	c := core.DefaultContext
	return fmt.Sprintf("%s:%s", l.Pkg.Get(c).String(), l.Name.Get(c))
}
