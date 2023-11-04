package build

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/AlexTheProgrammer/raft/coppier"
)

const distFolder string = "dist"

// build generates the wasm binary to be included in the build dist
func Build(assets embed.FS) {

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
	coppier.Assets(assets)

	// copy across user defined contexnt
	coppier.UserDefinedContent()

}
