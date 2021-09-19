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
	// starlark_tasks.T_EvalStarlark(ctx, labels.ParseLabel("//internal/starlark_tasks:test.bzl"))

	fmt.Println("Exists:", starlark_tasks.T_RuleExists(ctx, labels.ParseLabel("//examples/empty_rule:helloworld")))

	fmt.Println("Stats: ", task.GetGlobalStats(ctx))
}
