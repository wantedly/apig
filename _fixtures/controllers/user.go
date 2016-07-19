package controllers

import (
	"net/http"

	"github.com/USERNAME/APISERVER/db"
	"github.com/USERNAME/APISERVER/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	db := db.DBInstance(c)
	fields := c.DefaultQuery("fields", "*")
	var users []models.User
	db.Select(fields).Find(&users)
	c.JSON(200, users)
}

func GetUser(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	fields := c.DefaultQuery("fields", "*")
	var user models.User
	err := db.Select(fields).First(&user, id).Error

	if err != nil {
		content := gin.H{"error": "user with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}

	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	db := db.DBInstance(c)
	var user models.User
	c.Bind(&user)
	if db.Create(&user).Error != nil {
		content := gin.H{"error": "error occured"}
		c.JSON(500, content)
		return
	}
	c.JSON(201, user)
}

func UpdateUser(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	var user models.User
	if db.First(&user, id).Error != nil {
		content := gin.H{"error": "user with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.Bind(&user)
	db.Save(&user)
	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	db := db.DBInstance(c)
	id := c.Params.ByName("id")
	var user models.User
	if db.First(&user, id).Error != nil {
		content := gin.H{"error": "user with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	db.Delete(&user)
	c.Writer.WriteHeader(http.StatusNoContent)
}
