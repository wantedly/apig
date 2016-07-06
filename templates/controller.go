package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetModelnames(c *gin.Context) {
	db := db.DBInstance(c)
	fields := c.DefaultQuery("fields", "*")
	var modelnames []models.Modelname
	db.Select(fields).Find(&modelnames)
	c.JSON(200, modelnames)
}

func GetModelname(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	fields := c.DefaultQuery("fields", "*")
	var modelname models.Modelname
	err := db.Select(fields).First(&modelname, id).Error
	if err != nil {
		content := gin.H{"error": "modelname with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.JSON(200, modelname)
}

func CreateModelname(c *gin.Context) {
	db := db.DBInstance(c)
	var modelname models.Modelname
	c.Bind(&modelname)
	if db.Create(&modelname).Error != nil {
		content := gin.H{"error": "error occured"}
		c.JSON(500, content)
		return
	}
	c.JSON(201, modelname)
}

func UpdateModelname(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	var modelname models.Modelname
	if db.First(&modelname, id).Error != nil {
		content := gin.H{"error": "modelname with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.Bind(&modelname)
	db.Save(&modelname)
	c.JSON(200, modelname)
}

func DeleteModelname(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	var modelname models.Modelname
	if db.First(&modelname, id).Error != nil {
		content := gin.H{"error": "modelname with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	db.Delete(&modelname)
	c.Writer.WriteHeader(http.StatusNoContent)
}
