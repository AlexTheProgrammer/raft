package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/AlexTheProgrammer/raft/build"
	"github.com/AlexTheProgrammer/raft/server"
	"github.com/urfave/cli"
)

const distFolder string = "dist"

//go:embed assets/*
var assets embed.FS

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
	build.Build(assets)
	server.Serve(assets)
	return nil
}

func buildCommand(c *cli.Context) error {
	fmt.Println("Building Raft bundle...")
	build.Build(assets)
	return nil
}
