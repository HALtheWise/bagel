package refs

import (
	"fmt"
	"strings"

	"github.com/HALtheWise/bagel/lib/core"
)

// StringRef is used to refer to generic "strings" in various forms
type StringRef uint32

var StringTable core.InternTable[StringRef, string]

func (r StringRef) Get(c *core.Context) string {
	return StringTable.Get(c, r)
}

type LabelRef uint32
type Label struct {
	Pkg  StringRef
	Name StringRef
}

var LabelTable core.InternTable[LabelRef, Label]

func (r LabelRef) Get(c *core.Context) Label {
	return LabelTable.Get(c, r)
}

func (r StringRef) String() string {
	c := core.DefaultContext
	val := StringTable.Get(c, r)
	return fmt.Sprintf("ref(%s)", val)
}

func (r LabelRef) String() string {
	c := core.DefaultContext
	val := LabelTable.Get(c, r)
	return fmt.Sprintf("ref(%+v)", val)
}

func (l Label) String() string {
	c := core.DefaultContext
	return fmt.Sprintf("//%s:%s", StringTable.Get(c, l.Pkg), StringTable.Get(c, l.Name))
}

func ParseLabel(c *core.Context, label string) LabelRef {
	if !strings.HasPrefix(label, "//") {
		return core.INVALID
	}
	pkg, name, ok := strings.Cut(label[2:], ":")
	if !ok {
		return core.INVALID
	}
	return LabelTable.Insert(c, Label{
		StringTable.Insert(c, pkg),
		StringTable.Insert(c, name),
	})
}

func ParseRelativeLabel(c *core.Context, label string, from StringRef) LabelRef {
	if !strings.HasPrefix(label, ":") {
		return LabelTable.Insert(c, Label{from, StringTable.Insert(c, label[1:])})
	}
	return ParseLabel(c, label)
}
