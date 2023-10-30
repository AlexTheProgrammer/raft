package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "raft"
	app.Usage = "a golang SPA powered by wasm"

	app.Commands = []cli.Command{
		{
			Name:   "dev",
			Usage:  "Start Raft in development mode",
			Action: devCommand,
		},
		{
			Name:   "build",
			Usage:  "Generate single page application",
			Action: buildCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func devCommand(c *cli.Context) error {
	fmt.Println("Starting Raft in development mode...")
	// Implement your 'raft dev' functionality here.
	return nil
}

func buildCommand(c *cli.Context) error {
	fmt.Println("Building Raft bundle...")
	// Implement your 'raft build' functionality here.
	return nil
}
