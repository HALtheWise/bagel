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
	fmt.Printf("%+v  (%v)\n", loaded, loaded.Rule.Impl)

	analyzed := analysis.T_AnalyzeTarget(c, label)
	fmt.Println(analyzed.Providers)
}
