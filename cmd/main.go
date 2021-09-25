package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/HALtheWise/bagel/internal/refs"
)

var _ refs.GlobalContext

func query(c *cli.Context) error {
	if c.Bool("keep_going") {
		fmt.Println("Keeping going")
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:        "bagel",
		Description: "a lightweight implementation of Bazel written in Go",
		Commands: []*cli.Command{
			{
				Name:        "query",
				Description: "list unconfigured rules",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "keep_going"},
				},
				Action: query,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
