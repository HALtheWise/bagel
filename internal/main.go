package main

import "fmt"

func echo[T any](x T) T {
	return x
}

func main() {
	fmt.Println(echo("Hello World"))
}
