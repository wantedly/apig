package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gedex/inflector"
)

const templateDir = "_templates"

var funcMap = template.FuncMap{
	"pluralize": inflector.Pluralize,
	"tolower":   strings.ToLower,
}

var skeletons = []string{
	"README.md.tmpl",
	".gitignore.tmpl",
	"main.go.tmpl",
	filepath.Join("db", "db.go.tmpl"),
	filepath.Join("router", "router.go.tmpl"),
	filepath.Join("middleware", "set_db.go.tmpl"),
	filepath.Join("server", "server.go.tmpl"),
	filepath.Join("controllers", ".gitkeep.tmpl"),
	filepath.Join("models", ".gitkeep.tmpl"),
}

func generateSkeleton(detail Detail, outDir string) error {
	if fileExists(outDir) {
		fmt.Fprintf(os.Stderr, "%s is already exists", outDir)
		os.Exit(1)
	}

	for _, filename := range skeletons {
		srcPath := filepath.Join(templateDir, "skeleton", filename)
		dstPath := filepath.Join(outDir, strings.Replace(filename, ".tmpl", "", 1))

		body, err := Asset(srcPath)

		if err != nil {
			return err
		}

		tmpl, err := template.New("complex").Parse(string(body))

		if err != nil {
			return err
		}

		var buf bytes.Buffer

		if err := tmpl.Execute(&buf, detail); err != nil {
			return err
		}

		if !fileExists(filepath.Dir(dstPath)) {
			if err := mkdir(filepath.Dir(dstPath)); err != nil {
				return err
			}
		}

		if err := ioutil.WriteFile(dstPath, buf.Bytes(), 0644); err != nil {
			return err
		}
	}

	return nil
}

func generateController(model *Model, outDir string) error {
	body, err := Asset(filepath.Join(templateDir, "controller.go.tmpl"))

	if err != nil {
		return err
	}

	tmpl, err := template.New("controller").Funcs(funcMap).Parse(string(body))

	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, model); err != nil {
		return err
	}

	dstPath := filepath.Join(outDir, "controllers", strings.ToLower(model.Name)+".go")

	if !fileExists(filepath.Dir(dstPath)) {
		if err := mkdir(filepath.Dir(dstPath)); err != nil {
			return err
		}
	}

	if err := ioutil.WriteFile(dstPath, buf.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}

func generateREADME(models []*Model, outDir string) error {
	body, err := Asset(filepath.Join(templateDir, "README.md.tmpl"))

	if err != nil {
		return err
	}

	tmpl, err := template.New("readme").Funcs(funcMap).Parse(string(body))

	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, models); err != nil {
		return err
	}

	dstPath := filepath.Join(outDir, "README.md")

	if !fileExists(filepath.Dir(dstPath)) {
		if err := mkdir(filepath.Dir(dstPath)); err != nil {
			return err
		}
	}

	if err := ioutil.WriteFile(dstPath, buf.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}

func generateRouter(models []*Model, outDir string) error {
	body, err := Asset(filepath.Join(templateDir, "router.go.tmpl"))

	if err != nil {
		return err
	}

	tmpl, err := template.New("router").Funcs(funcMap).Parse(string(body))

	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, models); err != nil {
		return err
	}

	dstPath := filepath.Join(outDir, "router", "router.go")

	if !fileExists(filepath.Dir(dstPath)) {
		if err := mkdir(filepath.Dir(dstPath)); err != nil {
			return err
		}
	}

	if err := ioutil.WriteFile(dstPath, buf.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}
