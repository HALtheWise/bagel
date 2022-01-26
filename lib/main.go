package main

import (
	"fmt"
	"os"

	"github.com/HALtheWise/bagel/lib/core"
	"github.com/HALtheWise/bagel/lib/refs"
)

func main() {
	if workDir := os.Getenv("BUILD_WORKING_DIRECTORY"); workDir != "" {
		os.Chdir(workDir)
	}

	c := core.NewContext()

	label, err := refs.ParseLabel(c, os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(label)
}
