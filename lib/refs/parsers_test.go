package refs_test

import (
	"testing"

	"github.com/HALtheWise/bagel/lib/core"
	"github.com/HALtheWise/bagel/lib/refs"
)

func label(c *core.Context, repo, pkg, name string) refs.LabelRef {
	return refs.LabelTable.Insert(c,
		refs.Label{
			Pkg: refs.PackageTable.Insert(c, refs.Package{
				Workspace: refs.StringTable.Insert(c, repo),
				RelPath:   refs.StringTable.Insert(c, pkg),
			}),
			Name: refs.StringTable.Insert(c, name),
		})
}

func TestParser(t *testing.T) {
	testCases := []struct {
		desc            string
		label, from     string
		repo, pkg, name string
	}{
		{
			desc:  "Full",
			label: "@a//b/c:d", from: "@nowhere//:important",
			repo: "a", pkg: "b/c", name: "d",
		},
		{
			desc:  "NoRepo",
			label: "//b/c:d", from: "@a//:important",
			repo: "a", pkg: "b/c", name: "d",
		},
		{
			desc:  "NoPkg",
			label: ":d", from: "@a//b/c:e",
			repo: "a", pkg: "b/c", name: "d",
		},
		{
			desc:  "NoColon",
			label: "d", from: "@a//b/c:e",
			repo: "a", pkg: "b/c", name: "d",
		},
		{
			desc:  "ImplicitLabel",
			label: "@a//b/c", from: "@//",
			repo: "a", pkg: "b/c", name: "c",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			c := core.NewContext()
			want := label(c, tC.repo, tC.pkg, tC.name)
			from := refs.ParseLabel(c, tC.from, refs.INVALID_PACKAGE).Get(c).Pkg
			if got := refs.ParseLabel(c, tC.label, from); got != want {
				t.Errorf("Parsing %#v yielded %s, want %s", tC.label, got, want)
			}
		})
	}
}
