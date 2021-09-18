package main

import (
	"fmt"

	"github.com/HALtheWise/balez/internal/task"
)

var greetT = task.Task1("greet", func(ctx *task.Context, s string) string {
	return "Hello " + s
})

func main() {
	ctx := task.Root()
	fmt.Println(greetT(ctx, "Balez"))
	fmt.Println(greetT(ctx, "Balez"))

	fmt.Println(task.GetGlobalStats(ctx))
}
