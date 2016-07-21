package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func parseField(field *ast.Field) *Field {
	fieldNames := []string{}

	for _, name := range field.Names {
		fieldNames = append(fieldNames, name.Name)
	}

	var fieldType string
	var fieldTag string

	switch x := field.Type.(type) {
	case *ast.Ident: // e.g. string
		fieldType = x.Name

	case *ast.ArrayType: // e.g. []Email
		switch x2 := x.Elt.(type) {
		case *ast.Ident:
			fieldType = "[]" + x2.Name

		case *ast.StarExpr: // e.g. []*Email
			switch x3 := x2.X.(type) {
			case *ast.Ident:
				fieldType = "[]*" + x3.Name
			}
		}

	case *ast.StarExpr:
		switch x2 := x.X.(type) {
		case *ast.Ident: // e.g. *Profile
			fieldType = "*" + x2.Name

		case *ast.SelectorExpr: // e.g. *time.Time
			switch x3 := x2.X.(type) {
			case *ast.Ident:
				fieldType = "*" + x3.Name + "." + x2.Sel.Name
			}
		}
	}

	s, err := strconv.Unquote(field.Tag.Value)

	if err != nil {
		s = field.Tag.Value
	}

	jsonName := strings.Split((reflect.StructTag)(s).Get("json"), ",")[0]

	fieldTag = field.Tag.Value

	if len(fieldNames) != 1 {
		fmt.Fprintf(os.Stderr, "Failed to read model files. Please fix struct %s", fieldNames[0])
		os.Exit(1)
	}

	fs := Field{
		Name:     fieldNames[0],
		JSONName: jsonName,
		Type:     fieldType,
		Tag:      fieldTag,
	}
	return &fs
}

func parseModel(path string) ([]*Model, error) {
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
				var fields []*Field

				var modelName string

				switch x2 := spec.(type) {
				case *ast.TypeSpec:
					modelName = x2.Name.Name

					switch x3 := x2.Type.(type) {
					case *ast.StructType:
						for _, field := range x3.Fields.List {
							fs := parseField(field)
							fields = append(fields, fs)
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

func parseMain(path string) ([]string, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)

	if err != nil {
		return nil, err
	}

	var importPaths []string

	ast.Inspect(f, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.GenDecl:
			if x.Tok != token.IMPORT {
				break
			}

			for _, spec := range x.Specs {
				switch x2 := spec.(type) {
				case *ast.ImportSpec:
					importPaths = append(importPaths, strings.Trim(x2.Path.Value, "\""))
				}
			}
		}
		return true
	})
	return importPaths, nil
}
