package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/AlexTheProgrammer/raft/server"
	"github.com/urfave/cli"
)

const distFolder string = "dist"

//go:embed assets/*
var assets embed.FS

func copyAssets() {
	files := readFiles()

	for k, v := range files {
		err := os.WriteFile(filepath.Join(distFolder, k), v, 0644)
		if err != nil {
			log.Fatalf("could not create asset %q", k)
		}
	}
}

func copyUserDefinedContent() {
	targetFile := "styles.css"
	data, err := os.ReadFile(targetFile)
	if err != nil {
		log.Printf("no styles.css found, not including user defined css in build")
	}

	err = os.WriteFile(filepath.Join(distFolder, targetFile), data, 0644)
	if err != nil {
		log.Fatalf("could not create asset %q", targetFile)
	}
}

func readFiles() map[string][]byte {

	resp := make(map[string][]byte)
	files := []string{
		"index.html",
		"script.js",
		"wasm_exec.js",
	}

	// Read and display the file content
	for _, f := range files {

		fBytes, err := assets.ReadFile("assets/" + f)
		if err != nil {
			log.Fatalf("could not load in file %q error %v", f, err)
		}
		resp[f] = fBytes
	}

	return resp
}

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
	build()
	server.Serve()
	return nil
}

func buildCommand(c *cli.Context) error {
	fmt.Println("Building Raft bundle...")
	build()
	return nil
}

// build generates the wasm binary to be included in the build dist
func build() {
	// Command to build a Go program in the current directory
	os.Setenv("GOOS", "js")
	os.Setenv("GOARCH", "wasm")
	cmd := exec.Command("go", "build", "-o", filepath.Join(distFolder, "main.wasm"))

	cmd.Dir = "."

	// Redirect and display the command's output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the 'go build' command
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	// copy across assets
	copyAssets()

	// copy across user defined contexnt
	copyUserDefinedContent()

}
