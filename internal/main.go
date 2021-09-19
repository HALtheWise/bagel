package main

import (
	"fmt"
	"os"

	"github.com/HALtheWise/balez/internal/labels"
	"github.com/HALtheWise/balez/internal/starlark_tasks"
	"github.com/HALtheWise/balez/internal/task"
)

func main() {
	if workDir := os.Getenv("BUILD_WORKING_DIRECTORY"); workDir != "" {
		os.Chdir(workDir)
	}

	ctx := task.Root()

	label := labels.ParseLabel(os.Args[1])

	fmt.Println("Exists:", starlark_tasks.T_RuleExists(ctx, label))

	fmt.Println("Stats: ", task.GetGlobalStats(ctx))
}
