package main

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"
	"fmt"
	"os"

	"github.com/gedex/inflector"
)

const templateDir = "_templates"

var funcMap = template.FuncMap{
	"pluralize": inflector.Pluralize,
	"tolower":   strings.ToLower,
}

var staticFiles = []string{
	".gitignore",
	"main.go.tmpl",
	filepath.Join("db", "db.go"),
	filepath.Join("middleware", "set_db.go"),
	filepath.Join("server", "server.go.tmpl"),
	filepath.Join("controllers", ".gitkeep"),
	filepath.Join("models", ".gitkeep"),
}

func copyStaticFiles(importPath ImportPath, outDir string) error {
	if fileExists(outDir) {
		fmt.Println(outDir)
		fmt.Fprintf(os.Stderr, "%s is already exists", outDir)
		os.Exit(1)
	}

	for _, filename := range staticFiles {
		srcPath := filepath.Join(templateDir, filename)
		dstPath := filepath.Join(outDir, strings.TrimRight(filename, ".tmpl"))

		body, err := Asset(srcPath)

		if err != nil {
			return err
		}

		tmpl, err := template.New("complex").Funcs(funcMap).Parse(string(body))

		if err != nil {
			return err
		}

		var buf bytes.Buffer

		if err := tmpl.Execute(&buf, importPath); err != nil {
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
	body, err := Asset(filepath.Join(templateDir, "controllers", "controller.go.tmpl"))

	if err != nil {
		return err
	}

	tmpl, err := template.New("controller").Parse(string(body))

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
	body, err := Asset(filepath.Join(templateDir, "router", "router.go.tmpl"))

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
