package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: api-server-generator <model directory> <output directory>")
		os.Exit(1)
	}

	modelDir := os.Args[1]
	outDir := os.Args[2]

	if !fileExists(outDir) {
		if err := mkdir(outDir); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	files, err := ioutil.ReadDir(modelDir)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		modelPath := filepath.Join(modelDir, file.Name())
		models, err := parseFile(modelPath)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := generateRouter(models, outDir); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		for _, model := range models {
			if err := generateController(model, outDir); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	}

	if err := copyStaticFiles(outDir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
