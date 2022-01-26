package refs

import (
	"os"
	"path/filepath"

	"github.com/HALtheWise/bagel/lib/core"
)

var T_FilepathForLabel = core.MemoFunc1("T_FilepathForLabel",
	func(c *core.Context, label LabelRef) string {
		unpacked := label.Get(c)
		return filepath.Join(unpacked.Pkg.Get(c), unpacked.Name.Get(c))
	})

var T_FindBuildFile = core.MemoFunc1("T_FindBuildFile",
	func(c *core.Context, pkg StringRef) LabelRef {
		dir := pkg.Get(c)
		for _, buildfile := range []string{"BUILD.bagel", "BUILD.bazel", "BUILD"} {
			if _, err := os.Stat(filepath.Join(dir, buildfile)); err == nil {
				return LabelTable.Insert(c, Label{pkg, StringTable.Insert(c, buildfile)})
			}
		}
		return core.INVALID
	})
