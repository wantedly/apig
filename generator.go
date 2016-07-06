package main

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/gedex/inflector"
)

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

var routerTmpl = `
//{{ .Name }} API
		api.GET("/{{ pluralize (tolower .Name) }}", controllers.Get{{ pluralize .Name }})
		api.GET("/{{ pluralize (tolower .Name) }}/:id", controllers.Get{{ .Name }})
		api.POST("/{{ pluralize (tolower .Name) }}", controllers.Create{{ .Name }})
		api.PUT("/{{ pluralize (tolower .Name) }}/:id", controllers.Update{{ .Name }})
		api.DELETE("/{{ pluralize (tolower .Name) }}/:id", controllers.Delete{{ .Name }})
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

func generateRouter(model *Model) (string, error) {
	tmpl, err := template.New("router").Funcs(funcMap).Parse(routerTmpl)

	if err != nil {
		return "", nil
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, model); err != nil {
		return "", nil
	}

	return strings.TrimSpace(buf.String()), nil
}
