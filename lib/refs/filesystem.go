package refs

import (
	"os"
	"path/filepath"

	"github.com/HALtheWise/bagel/lib/core"
)

var emptystring = StringTable.Insert(nil, "")

func FilepathForPackage(c *core.Context, pkg PackageRef) string {
	unpacked := pkg.Get(c)

	workspace := ""
	if unpacked.Workspace != emptystring {
		workspace = "external/" + unpacked.Workspace.Get(c)
	}

	return filepath.Join(workspace, unpacked.RelPath.Get(c))
}

var T_FilepathForLabel = core.Task1("T_FilepathForLabel",
	func(c *core.Context, label LabelRef) string {
		unpacked := label.Get(c)

		return filepath.Join(FilepathForPackage(c, unpacked.Pkg), unpacked.Name.Get(c))
	})

var T_FindBuildFile = core.Task1("T_FindBuildFile",
	func(c *core.Context, pkg PackageRef) LabelRef {
		dir := FilepathForPackage(c, pkg)
		for _, buildfile := range []string{"BUILD.bagel", "BUILD.bazel", "BUILD"} {
			if _, err := os.Stat(filepath.Join(dir, buildfile)); err == nil {
				return LabelTable.Insert(c, Label{pkg, StringTable.Insert(c, buildfile)})
			}
		}
		return INVALID_LABEL
	})
