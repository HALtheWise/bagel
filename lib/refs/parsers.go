package refs

import (
	"fmt"
	"strings"

	"github.com/HALtheWise/bagel/lib/core"
)

func ParsePackage(c *core.Context, s string, from PackageRef) PackageRef {
	workspace, relpath, ok := strings.Cut(s, "//")
	if !ok {
		panic("Not a package: " + s)
		return INVALID_PACKAGE
	}
	var workspaceRef StringRef
	if len(workspace) == 0 {
		workspaceRef = from.Get(c).Workspace
	} else {
		if !strings.HasPrefix(workspace, "@") {
			panic("Not a package: " + s)
			return INVALID_PACKAGE
		}
		workspaceRef = StringTable.Insert(c, workspace[1:])
	}
	return PackageTable.Insert(c,
		Package{
			Workspace: workspaceRef,
			RelPath:   StringTable.Insert(c, relpath),
		})
}

func ParseLabel(c *core.Context, label string, from PackageRef) LabelRef {
	pkg, name, ok := strings.Cut(label, ":")

	if !ok {
		if strings.Contains(pkg, "//") {
			// This is of the form //a/b (short for //a/b:b)
			name = pkg[strings.LastIndex(pkg, "/")+1:]
		} else {
			// this is of the form file.cc (short for :file.cc)
			name = pkg
			pkg = ""
		}
	}

	if len(pkg) == 0 {
		// This is a relative label
		if from == INVALID_PACKAGE {
			panic(fmt.Sprintf("Relative label %s encountered in invalid context", label))
		}
		return LabelTable.Insert(c, Label{from, StringTable.Insert(c, name)})
	}
	return LabelTable.Insert(c, Label{
		ParsePackage(c, pkg, from),
		StringTable.Insert(c, name),
	})
}
