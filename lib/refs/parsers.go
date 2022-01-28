package refs

import (
	"strings"

	"github.com/HALtheWise/bagel/lib/core"
)

func ParsePackage(c *core.Context, s string) PackageRef {
	workspace, relpath, ok := strings.Cut(s, "//")
	if !ok {
		panic("Not a package: " + s)
		return core.INVALID
	}
	if len(workspace) > 0 {
		if !strings.HasPrefix(workspace, "@") {
			panic("Not a package: " + s)
			return core.INVALID
		}
		workspace = workspace[1:]
	}
	return PackageTable.Insert(c,
		Package{
			Workspace: StringTable.Insert(c, workspace),
			RelPath:   StringTable.Insert(c, relpath),
		})
}

func ParseLabel(c *core.Context, label string) LabelRef {
	pkg, name, ok := strings.Cut(label, ":")
	if !ok {
		panic("Invalid label: " + label)
		return core.INVALID
	}
	return LabelTable.Insert(c, Label{
		ParsePackage(c, pkg),
		StringTable.Insert(c, name),
	})
}

func ParseRelativeLabel(c *core.Context, label string, from PackageRef) LabelRef {
	if strings.HasPrefix(label, ":") {
		return LabelTable.Insert(c, Label{from, StringTable.Insert(c, label[1:])})
	}
	return ParseLabel(c, label)
}
