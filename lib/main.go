package main

import (
	"fmt"
	"os"

	"github.com/HALtheWise/bagel/lib/analysis"
	"github.com/HALtheWise/bagel/lib/core"
	"github.com/HALtheWise/bagel/lib/loading"
	"github.com/HALtheWise/bagel/lib/refs"
)

func main() {
	if workDir := os.Getenv("BUILD_WORKING_DIRECTORY"); workDir != "" {
		os.Chdir(workDir)
	}

	c := core.NewContext()

	label := refs.ParseLabel(c, os.Args[1])

	fmt.Println(label)

	loaded := loading.T_LoadTarget(c, label)
	fmt.Printf("%+v  (%v)\n\n", loaded, loaded.Rule.Impl)

	clabel := refs.CLabelTable.Insert(c, refs.CLabel{Label: label, Config: refs.TARGET_CONFIG})

	analyzed := analysis.T_AnalyzeTarget(c, clabel)
	fmt.Printf("AnalyzedTarget: %+v\n\n", analyzed)

	file := analyzed.Providers[0].Data["files"].(*loading.Depset).Items[0].(*analysis.BzlFile)

	fileInfo := analysis.T_FileInfo(c, file.Ref)
	fmt.Printf("FileInfo: %+v\n\n", fileInfo)

	actionInfo := analysis.T_ActionInfo(c, fileInfo.Generator)
	fmt.Printf("ActionInfo: %+v\n\n", actionInfo)
}
