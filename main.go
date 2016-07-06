package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Model struct {
	Name   string
	Fields map[string]string
}

func parseFile(path string) ([]*Model, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)

	if err != nil {
		return nil, err
	}

	models := []*Model{}

	ast.Inspect(f, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.GenDecl:
			if x.Tok != token.TYPE {
				break
			}

			for _, spec := range x.Specs {
				fieldNames := []string{}
				fields := make(map[string]string)

				var modelName string

				switch x2 := spec.(type) {
				case *ast.TypeSpec:
					modelName = x2.Name.Name

					switch x3 := x2.Type.(type) {
					case *ast.StructType:
						for _, field := range x3.Fields.List {
							for _, name := range field.Names {
								fieldNames = append(fieldNames, name.Name)
							}

							var fieldType string

							switch x4 := field.Type.(type) {
							case *ast.Ident: // e.g. string
								fieldType = x4.Name
							case *ast.StarExpr: // e.g. *time.Time
								switch x5 := x4.X.(type) {
								case *ast.SelectorExpr:
									switch x6 := x5.X.(type) {
									case *ast.Ident:
										fieldType = x6.Name + "." + x5.Sel.Name
									}
								}
							}

							for _, name := range fieldNames {
								fields[name] = fieldType
							}
						}
					}

					models = append(models, &Model{
						Name:   modelName,
						Fields: fields,
					})
				}
			}
		}

		return true
	})

	return models, nil
}

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

		for _, model := range models {
			if err := generateController(model, outDir); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			if err := generateRouter(model, outDir); err != nil {
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
