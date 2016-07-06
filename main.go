package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gedex/inflector"
)

type Model struct {
	Name   string
	Fields map[string]string
}

var controllerTmpl = `
func Get{{ pluralize .Name }}(c *gin.Context) {
	db := db.DBInstance(c)
	fields := c.DefaultQuery("fields", "*")
	var {{ pluralize (tolower .Name) }} []models.{{ .Name }}
	db.Select(fields).Find(&{{ pluralize (tolower .Name) }})
	c.JSON(200, {{ pluralize (tolower .Name) }})
}

func Get{{ .Name }}(c. *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	fields := c.DefaultQuery("fields", "*")
	var {{ tolower .Name }} models.{{ .Name }}
	err := db.Select(fields).First(&{{ tolower .Name }}, id).Error

	if err != nil {
		content := gin.H{"error": "{{ tolower .Name }} with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}

	c.JSON(200, {{ tolower .Name }})
}

func Create{{ .Name }}(c *gin.Context) {
	db := db.DBInstance(c)
	var {{ tolower .Name }} models.{{ .Name }}
	c.Bind(&{{ tolower .Name }})
	if db.Create(&{{ tolower .Name }}).Error != nil {
		content := gin.H{"error": "error occured"}
		c.JSON(500, content)
		return
	}
	c.JSON(201, {{ tolower .Name }})
}

func Update{{ .Name }}(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	var {{ tolower .Name }} models.{{ .Name }}
	if db.First(&{{ tolower .Name }}, id).Error != nil {
		content := gin.H{"error": "{{ tolower .Name }} with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.Bind(&{{ tolower .Name }})
	db.Save(&{{ tolower .Name }})
	c.JSON(200, {{ tolower .Name }})
}

func Delete{{ .Name }}(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	var {{ tolower .Name }} models.{{ .Name }}
	if db.First(&{{ tolower .Name }}, id).Error != nil {
		content := gin.H{"error": "{{ tolower .Name }} with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	db.Delete(&{{ tolower .Name }})
	c.Writer.WriteHeader(http.StatusNoContent)
}
`

var funcMap = template.FuncMap{
	"pluralize": inflector.Pluralize,
	"tolower":   strings.ToLower,
}

func generateController(model *Model) (string, error) {
	tmpl, err := template.New("controller").Funcs(funcMap).Parse(controllerTmpl)

	if err != nil {
		return "", nil
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, model); err != nil {
		return "", nil
	}

	return strings.TrimSpace(buf.String()), nil
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
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: api-server-generator <model directory>")
		os.Exit(1)
	}

	dir := os.Args[1]

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		path := filepath.Join(dir, file.Name())
		fmt.Println("===== " + path)

		models, err := parseFile(path)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		for _, model := range models {
			fmt.Println("// Model: " + model.Name)
			fmt.Println("// Fields: ")

			for name, t := range model.Fields {
				fmt.Println("//  - " + name + " => " + t)
			}

			fmt.Println("")

			controller, err := generateController(model)

			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			fmt.Println(controller)
		}
	}
}
