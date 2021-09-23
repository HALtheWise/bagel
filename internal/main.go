package main

import (
	"fmt"
	"os"

	"github.com/HALtheWise/bales/internal/labels"
	"github.com/HALtheWise/bales/internal/starlark_tasks"
	"github.com/HALtheWise/bales/internal/task"
)

func main() {
	if workDir := os.Getenv("BUILD_WORKING_DIRECTORY"); workDir != "" {
		os.Chdir(workDir)
	}

	ctx := task.Root()

	label := labels.ParseLabel(os.Args[1])

	fmt.Println("Unconfigured:", starlark_tasks.T_RuleInfoUnconfigured(ctx, label))

	fmt.Println("Configured:", starlark_tasks.T_RuleInfoEvaluated(ctx, label))

	fmt.Println("Stats: ", task.GetGlobalStats(ctx))
}
