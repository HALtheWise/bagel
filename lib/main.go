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

	info := loading.T_RuleInfoUnconfigured(c, label)
	fmt.Println(info.Impl)

	rule := analysis.T_RuleInfoEvaluated(c, label)
	fmt.Println(rule.Providers)
}
