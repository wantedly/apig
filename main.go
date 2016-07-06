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

		generateController(path)
	}
}

func generateController(path string) {
	structName, err := modelName(path)
	ifErr(err)
	fmt.Println("ModelName: " + structName)
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
