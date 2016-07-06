package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./api-server-generator <modles directory>")
		os.Exit(1)
	}

	dir := os.Args[1]

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join(dir, file.Name())
		fmt.Println(filePath)
	}
}
