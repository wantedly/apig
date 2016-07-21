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
	"apibDefaultValue": apibDefaultValue,
	"apibType":         apibType,
	"pluralize":        inflector.Pluralize,
	"requestParams":    requestParams,
	"tolower":          strings.ToLower,
	"title":            strings.Title,
}

var skeletons = []string{
	"README.md.tmpl",
	".gitignore.tmpl",
	"main.go.tmpl",
	filepath.Join("db", "db.go.tmpl"),
	filepath.Join("router", "router.go.tmpl"),
	filepath.Join("middleware", "set_db.go.tmpl"),
	filepath.Join("server", "server.go.tmpl"),
	filepath.Join("version", "version.go.tmpl"),
	filepath.Join("version", "version_test.go.tmpl"),
	filepath.Join("controllers", ".gitkeep.tmpl"),
	filepath.Join("docs", ".gitkeep.tmpl"),
	filepath.Join("models", ".gitkeep.tmpl"),
}

var managedFields = []string{
	"ID",
	"CreatedAt",
	"UpdatedAt",
}

func apibDefaultValue(field *ModelField) string {
	switch field.Type {
	case "bool":
		return "false"
	case "string":
		return strings.ToUpper(field.Name)
	case "time.Time":
		return "`2000-01-01 00:00:00`"
	case "uint":
		return "1"
	}

	return strings.ToUpper(field.Name)
}

func apibType(field *ModelField) string {
	switch field.Type {
	case "bool":
		return "boolean"
	case "string":
		return "string"
	case "uint":
		return "number"
	}

	return "string"
}

func requestParams(fields []*ModelField) []*ModelField {
	var managed bool

	params := []*ModelField{}

	for _, field := range fields {
		managed = false

		for _, name := range managedFields {
			if field.Name == name {
				managed = true
			}
		}

		if !managed {
			params = append(params, field)
		}
	}

	return params
}

func generateApibIndex(detail *Detail, outDir string) error {
	body, err := Asset(filepath.Join(templateDir, "index.apib.tmpl"))

	if err != nil {
		return err
	}

	tmpl, err := template.New("apib").Funcs(funcMap).Parse(string(body))

	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, detail); err != nil {
		return err
	}

	dstPath := filepath.Join(outDir, "docs", "index.apib")

	if !fileExists(filepath.Dir(dstPath)) {
		if err := mkdir(filepath.Dir(dstPath)); err != nil {
			return err
		}
	}

	if err := ioutil.WriteFile(dstPath, buf.Bytes(), 0644); err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "\t\x1b[32m%s\x1b[0m %s\n", "create", dstPath)

	return nil
}

func generateApibModel(detail *Detail, outDir string) error {
	body, err := Asset(filepath.Join(templateDir, "model.apib.tmpl"))

	if err != nil {
		return err
	}

	tmpl, err := template.New("apib").Funcs(funcMap).Parse(string(body))

	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, detail); err != nil {
		return err
	}

	dstPath := filepath.Join(outDir, "docs", strings.ToLower(detail.Model.Name)+".apib")

	if !fileExists(filepath.Dir(dstPath)) {
		if err := mkdir(filepath.Dir(dstPath)); err != nil {
			return err
		}
	}

	if err := ioutil.WriteFile(dstPath, buf.Bytes(), 0644); err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "\t\x1b[32m%s\x1b[0m %s\n", "create", dstPath)

	return nil
}

func generateSkeleton(detail *Detail, outDir string) error {
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

		fmt.Printf("\t\x1b[32m%s\x1b[0m %s\n", "create", dstPath)
	}

	return nil
}

func generateController(detail *Detail, outDir string) error {
	body, err := Asset(filepath.Join(templateDir, "controller.go.tmpl"))

	if err != nil {
		return err
	}

	tmpl, err := template.New("controller").Funcs(funcMap).Parse(string(body))

	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, detail); err != nil {
		return err
	}

	dstPath := filepath.Join(outDir, "controllers", strings.ToLower(detail.Model.Name)+".go")

	if !fileExists(filepath.Dir(dstPath)) {
		if err := mkdir(filepath.Dir(dstPath)); err != nil {
			return err
		}
	}

	if err := ioutil.WriteFile(dstPath, buf.Bytes(), 0644); err != nil {
		return err
	}

	fmt.Printf("\t\x1b[32m%s\x1b[0m %s\n", "create", dstPath)

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

	fmt.Fprintf(os.Stdout, "\t\x1b[32m%s\x1b[0m %s\n", "update", dstPath)

	return nil
}

func generateRouter(detail *Detail, outDir string) error {
	body, err := Asset(filepath.Join(templateDir, "router.go.tmpl"))

	if err != nil {
		return err
	}

	tmpl, err := template.New("router").Funcs(funcMap).Parse(string(body))

	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, detail); err != nil {
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

	fmt.Fprintf(os.Stdout, "\t\x1b[32m%s\x1b[0m %s\n", "update", dstPath)

	return nil
}
