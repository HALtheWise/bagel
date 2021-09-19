package main

import (
	"os"

	"github.com/HALtheWise/balez/internal/starlark_tasks"
	"github.com/HALtheWise/balez/internal/task"
)

func main() {
	if workDir := os.Getenv("BUILD_WORKING_DIRECTORY"); workDir != "" {
		os.Chdir(workDir)
	}

	ctx := task.Root()
	starlark_tasks.T_ExecuteFile(ctx, "internal/starlark_tasks/test.bzl")
}
