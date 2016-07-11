package main

import (
	"bytes"
	"io/ioutil"
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

var staticFiles = []string{
	".gitignore",
	"main.go",
	filepath.Join("db", "db.go"),
	filepath.Join("middleware", "set_db.go"),
	filepath.Join("server", "server.go"),
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

		body, err := Asset(srcPath)

		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(dstPath, body, 0644); err != nil {
			return err
		}
	}

	return nil
}

func generateApib(model *Model, outDir string) error {
	body, err := Asset(filepath.Join(templateDir, "apib.apib.tmpl"))

	if err != nil {
		return err
	}

	tmpl, err := template.New("apib").Funcs(funcMap).Parse(string(body))

	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, model); err != nil {
		return err
	}

	dstPath := filepath.Join(outDir, strings.ToLower(model.Name)+".apib")

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

func generateController(model *Model, outDir string) error {
	body, err := Asset(filepath.Join(templateDir, "controllers", "controller.go.tmpl"))

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
