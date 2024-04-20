package main

import (
	"fmt"
	"os"

	"github.com/HALtheWise/bagel/lib/core"
	"github.com/HALtheWise/bagel/lib/execution"
	"github.com/HALtheWise/bagel/lib/refs"
)

func main() {
	if workDir := os.Getenv("BUILD_WORKING_DIRECTORY"); workDir != "" {
		os.Chdir(workDir)
	}

	c := core.NewContext()

	from := refs.PackageTable.Insert(c, refs.Package{
		Workspace: refs.EMPTYSTRING,
		RelPath:   refs.EMPTYSTRING,
	})
	label := refs.ParseLabel(c, os.Args[1], from)

	fmt.Println("Building DefaultInfo for:", label)

	err := execution.T_BuildDefaultInfo(c, label)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
