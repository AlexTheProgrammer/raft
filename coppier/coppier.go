package coppier

import (
	"embed"
	"log"
	"os"
	"path/filepath"
)

const distFolder string = "dist"

func UserDefinedContent() {

	log.Printf("copying across user defined content")

	targetFile := "styles.css"
	data, err := os.ReadFile(targetFile)
	if err != nil {
		log.Printf("no styles.css found, not including user defined css in build")
	}

	outputPath := filepath.Join(distFolder, targetFile)
	err = os.WriteFile(outputPath, data, 0644)
	if err != nil {
		log.Fatalf("could not create asset %q", targetFile)
	}

	log.Printf("user defined content %v copied to %v", targetFile, outputPath)
}

func Assets(assets embed.FS) {
	files := readFiles(assets)

	for k, v := range files {
		err := os.WriteFile(filepath.Join(distFolder, k), v, 0644)
		if err != nil {
			log.Fatalf("could not create asset %q", k)
		}
	}
}

func readFiles(assets embed.FS) map[string][]byte {

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
