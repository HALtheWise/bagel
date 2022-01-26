package main

import (
	"fmt"
	"os"

	"github.com/HALtheWise/bagel/lib/core"
	"github.com/HALtheWise/bagel/lib/refs"
	"github.com/HALtheWise/bagel/lib/starlark_tasks"
)

func main() {
	if workDir := os.Getenv("BUILD_WORKING_DIRECTORY"); workDir != "" {
		os.Chdir(workDir)
	}

	c := core.NewContext()

	label := refs.ParseLabel(c, os.Args[1])

	fmt.Println(label)

	info := starlark_tasks.T_RuleInfoUnconfigured(c, label)
	fmt.Println(info.Impl)

	rule := starlark_tasks.T_RuleInfoEvaluated(c, label)
	fmt.Println(rule.Providers)
}
