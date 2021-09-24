package labels

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/HALtheWise/bagel/internal/task"
)

type Package struct {
	Repo, Pkg string
}

type Label struct {
	Package
	Name string
}

var NullPackage = Package{"__NULL", ""}
var NullLabel = Label{NullPackage, ""}

var T_FindBuildFile = task.Task1("T_FindBuildFile", func(c *task.Context, pkg Package) Label {
	for _, name := range []string{"BUILD.bagel", "BUILD.bazel", "BUILD"} {
		l := Label{pkg, name}
		if info, err := os.Stat(T_FilepathForLabel(c, l)); err == nil && info.Mode().IsRegular() {
			return l
		}
	}
	return NullLabel
})

// TODO(eric): This should maybe _not_ be a tracked task so it's not part of cache keys?
var T_FilepathForLabel = task.Task1("T_FilepathForLabel", func(c *task.Context, l Label) string {
	if l == NullLabel {
		panic("Null label")
	}
	if l.Repo != "" {
		panic("We don't support repos yet, sorry!")
	}
	return filepath.Join(l.Pkg, l.Name)
})

func (l *Label) String() string {
	if l.Repo != "" {
		panic("We don't support repos yet, sorry!")
	}
	return fmt.Sprintf("//%s:%s", l.Pkg, l.Name)
}

func ParsePackage(s string) Package {
	if !strings.HasPrefix(s, "//") || strings.Contains(s, ":") {
		panic(fmt.Errorf("Malformed package name %s", s))
	}
	return Package{Repo: "", Pkg: s[2:]}
}

func ParseRelativeLabel(s string, from Package) Label {
	if strings.Contains(s, ":") && !strings.HasPrefix(s, ":") {
		// This is an absolute label
		return ParseLabel(s)
	}
	return Label{Package: from, Name: strings.TrimPrefix(s, ":")}
}

func ParseLabel(s string) Label {
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		panic("Label must have :")
	}
	return Label{Package: ParsePackage(parts[0]), Name: parts[1]}
}
