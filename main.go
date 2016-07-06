package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	ControllerTemp = "templates/controller.go"
	ControllerDir  = "controllers"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./api-server-generator <modles directory>")
		os.Exit(1)
	}

	dir := os.Args[1]

	files, err := ioutil.ReadDir(dir)
	ifErr(err)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		path := filepath.Join(dir, file.Name())
		fmt.Println(path)

		ifErr(generateController(path))
	}
}

func generateController(path string) error {
	structName, err := modelName(path)
	ifErr(err)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, ControllerTemp, nil, parser.ParseComments)
	ifErr(err)

	ifErr(rewriteTemplate(f, structName))

	_ = os.Mkdir(ControllerDir, 0755)

	filename := fmt.Sprintf("%s.go", strings.ToLower(structName))

	genPath := filepath.Join(ControllerDir, filename)

	file, err := os.Create(genPath)
	ifErr(err)

	defer file.Close()
	format.Node(file, fset, f)

	fmt.Println("  ===> " + genPath)

	return nil
}

func rewriteTemplate(f ast.Node, structName string) error {
	ast.Inspect(f, func(n ast.Node) bool {
		switch aType := n.(type) {
		case *ast.Ident:
			if strings.Contains(aType.Name, "Modelname") {
				aType.Name = strings.Replace(aType.Name, "Modelname", structName, 1)
			} else if strings.Contains(aType.Name, "modelname") {
				aType.Name = strings.Replace(aType.Name, "modelname", strings.ToLower(structName), 1)
			}

		case *ast.BasicLit:
			if strings.Contains(aType.Value, "modelname") {
				aType.Value = strings.Replace(aType.Value, "modelname", strings.ToLower(structName), 1)
			}
		}

		return true
	})
	return nil
}

func modelName(path string) (string, error) {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	ifErr(err)

	var structName string

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GenDecl:
			if x.Tok != token.TYPE {
				break
			}
			for _, spec := range x.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				_, ok = typeSpec.Type.(*ast.StructType)
				if !ok {
					continue
				}

				structName = typeSpec.Name.Name
			}
		}
		return true
	})
	return structName, nil
}

func ifErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
