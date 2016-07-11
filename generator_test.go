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

func TestGenerateController(t *testing.T) {
	model := &Model{
		Name: "User",
		Fields: map[string]string{
			"ID": "hoge",
		},
	}

	outDir, err := ioutil.TempDir("", "generateController")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	defer os.RemoveAll(outDir)

	if err := generateController(model, outDir); err != nil {
		t.Fatalf("Error should not be raised: %#v", err)
	}

	path := filepath.Join(outDir, "controllers", "user.go")
	_, err = os.Stat(path)
	if err != nil {
		t.Fatalf("Controller file is not generated: %s", path)
	}
}

func TestGenerateREADME(t *testing.T) {
	model := &Model{
		Name: "User",
		Fields: map[string]string{
			"ID": "hoge",
		},
	}
	models := []*Model{model}

	outDir, err := ioutil.TempDir("", "generateREADME")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	defer os.RemoveAll(outDir)

	if err := generateREADME(models, outDir); err != nil {
		t.Fatalf("Error should not be raised: %#v", err)
	}

	path := filepath.Join(outDir, "README.md")
	_, err = os.Stat(path)
	if err != nil {
		t.Fatalf("README is not generated: %s", path)
	}
}

func TestGenerateRouter(t *testing.T) {
	model := &Model{
		Name: "User",
		Fields: map[string]string{
			"ID": "hoge",
		},
	}
	models := []*Model{model}

	outDir, err := ioutil.TempDir("", "generateRouter")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	defer os.RemoveAll(outDir)

	if err := generateRouter(models, outDir); err != nil {
		t.Fatalf("Error should not be raised: %#v", err)
	}

	path := filepath.Join(outDir, "router", "router.go")
	_, err = os.Stat(path)
	if err != nil {
		t.Fatalf("Router file is not generated: %s", path)
	}
}
