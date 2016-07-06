package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gedex/inflector"
)

const templateDir = "templates"

var funcMap = template.FuncMap{
	"pluralize": inflector.Pluralize,
	"tolower":   strings.ToLower,
}

var staticFiles = []string{
	".gitignore",
	"README.md",
	"main.go",
	// filepath.Join("db", "db.go"),
	// filepath.Join("middleware", "db.go"),
	// filepath.Join("server", "server.go"),
}

func copyStaticFiles(outDir string) error {
	for _, filename := range staticFiles {
		srcPath := filepath.Join(templateDir, filename)
		dstPath := filepath.Join(outDir, filename)

		if !fileExists(filepath.Dir(dstPath)) {
			if err := mkdir(filepath.Dir(dstPath)); err != nil {
				return err
			}
		}

		src, err := os.Open(srcPath)

		if err != nil {
			return err
		}

		defer src.Close()

		dst, err := os.Create(dstPath)

		if err != nil {
			return err

		}
		defer dst.Close()

		_, err = io.Copy(dst, src)

		if err != nil {
			return err
		}
	}

	return nil
}

func generateController(model *Model, outDir string) error {
	tmpl, err := template.New("controller").Funcs(funcMap).Parse(controllerTmpl)

	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, model); err != nil {
		return err
	}

	dstPath := filepath.Join(outDir, "controller", strings.ToLower(model.Name)+".go")

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

func generateRouter(model *Model, outDir string) error {
	tmpl, err := template.New("router").Funcs(funcMap).Parse(routerTmpl)

	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, model); err != nil {
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
