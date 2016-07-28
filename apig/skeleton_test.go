package apig

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateSkeleton(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "copyStaticFiles")
	if err != nil {
		t.Fatal("Failed to create tempdir")
	}
	defer os.RemoveAll(tempDir)

	outDir := filepath.Join(tempDir, "api-server")

	if err := generateSkeleton(detail, outDir); err != nil {
		t.Fatalf("Error should not be raised: %s", err)
	}

	files := []string{
		"README.md",
		".gitignore",
		"main.go",
		filepath.Join("db", "db.go"),
		filepath.Join("db", "pagination.go"),
		filepath.Join("router", "router.go"),
		filepath.Join("middleware", "set_db.go"),
		filepath.Join("server", "server.go"),
		filepath.Join("helper", "field.go"),
		filepath.Join("helper", "field_test.go"),
		filepath.Join("version", "version.go"),
		filepath.Join("version", "version_test.go"),
		filepath.Join("controllers", ".gitkeep"),
		filepath.Join("models", ".gitkeep"),
	}

	for _, file := range files {
		_, err := os.Stat(filepath.Join(outDir, file))
		if err != nil {
			t.Fatalf("Static file is not copied: %s", file)
		}
	}
}
