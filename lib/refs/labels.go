package refs

import (
	"fmt"

	"github.com/HALtheWise/bagel/lib/core"
)

// StringRef is used to refer to generic "strings" in various forms
type StringRef uint32

var StringTable core.InternTable[StringRef, string]

func (r StringRef) Get(c *core.Context) string {
	return StringTable.Get(c, r)
}

type PackageRef uint32

type Package struct {
	Workspace StringRef
	RelPath   StringRef
}

var PackageTable core.InternTable[PackageRef, Package]

func (r PackageRef) Get(c *core.Context) Package {
	return PackageTable.Get(c, r)
}

type LabelRef uint32
type Label struct {
	Pkg  PackageRef
	Name StringRef
}

var LabelTable core.InternTable[LabelRef, Label]

func (r LabelRef) Get(c *core.Context) Label {
	return LabelTable.Get(c, r)
}

func (r StringRef) String() string {
	c := core.DefaultContext
	val := r.Get(c)
	return fmt.Sprintf(`r"%s"`, val)
}

func (r PackageRef) String() string {
	c := core.DefaultContext
	val := r.Get(c)
	return fmt.Sprintf("r%+v", val)
}
func (p Package) String() string {
	c := core.DefaultContext
	return fmt.Sprintf("@%s//%s", p.Workspace.Get(c), p.RelPath.Get(c))
}

func (r LabelRef) String() string {
	c := core.DefaultContext
	val := r.Get(c)
	return fmt.Sprintf("r%+v", val)
}

func (l Label) String() string {
	c := core.DefaultContext
	return fmt.Sprintf("%s:%s", l.Pkg.Get(c).String(), l.Name.Get(c))
}
