package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCopyStaticFiles(t *testing.T) {
	outDir, err := ioutil.TempDir("", "copyStaticFiles")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	defer os.RemoveAll(outDir)

	if err := copyStaticFiles(outDir); err != nil {
		t.Fatalf("Error should not be raised: %#v", err)
	}

	files := []string{
		".gitignore",
		"main.go",
		filepath.Join("db", "db.go"),
		filepath.Join("middleware", "set_db.go"),
		filepath.Join("server", "server.go"),
	}

	for _, file := range files {
		_, err := os.Stat(filepath.Join(outDir, file))
		if err != nil {
			t.Fatalf("Static file is not copied: %s", file)
		}
	}
}
