package controllers

import (
	"net/http"

	dbpkg "github.com/wantedly/apig/_example/db"
	"github.com/wantedly/apig/_example/helper"
	"github.com/wantedly/apig/_example/models"
	"github.com/wantedly/apig/_example/version"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func setUserPreload(fields []string, db *gorm.DB) ([]string, *gorm.DB) {
	sel := make([]string, len(fields))
	copy(sel, fields)
	offset := 0
	for key, val := range fields {
		switch val {

		case "jobs":
			db = db.Preload("Jobs")
			sel = append(sel[:(key-offset)], sel[(key+1-offset):]...)
			offset += 1
			idflag := true
			for _, v := range sel {
				if v == "id" {
					idflag = false
					break
				}
			}
			if idflag {
				sel = append(sel, "id")
			}

		case "emails":
			db = db.Preload("Emails")
			sel = append(sel[:(key-offset)], sel[(key+1-offset):]...)
			offset += 1
			idflag := true
			for _, v := range sel {
				if v == "id" {
					idflag = false
					break
				}
			}
			if idflag {
				sel = append(sel, "id")
			}

		}
	}
	return sel, db
}

func GetUsers(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pagination := dbpkg.Pagination{}
	db, err := pagination.Paginate(c)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fields, nestFields := helper.ParseFields(c.DefaultQuery("fields", "*"))
	sel, db := setUserPreload(fields, db)
	var users []models.User
	err = db.Select(sel).Find(&users).Error

	if err != nil {
		c.JSON(500, gin.H{"error": "error occured"})
		return
	}

	// paging
	var index int
	if len(users) < 1 {
		index = 0
	} else {
		index = int(users[len(users)-1].ID)
	}
	pagination.SetHeaderLink(c, index)

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	var fieldMap []map[string]interface{}
	for key, _ := range users {
		fieldMap = append(fieldMap, helper.FieldToMap(users[key], fields, nestFields))
	}
	c.JSON(200, fieldMap)
}

func GetUser(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	fields, nestFields := helper.ParseFields(c.DefaultQuery("fields", "*"))
	sel, db := setUserPreload(fields, db)
	var user models.User

	if db.Select(sel).First(&user, id).Error != nil {
		content := gin.H{"error": "user with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	fieldMap := helper.FieldToMap(user, fields, nestFields)
	c.JSON(200, fieldMap)
}

func CreateUser(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	var user models.User
	c.Bind(&user)
	if db.Create(&user).Error != nil {
		content := gin.H{"error": "error occured"}
		c.JSON(500, content)
		return
	}

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(201, user)
}

func UpdateUser(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	var user models.User
	if db.First(&user, id).Error != nil {
		content := gin.H{"error": "user with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.Bind(&user)
	db.Save(&user)

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	var user models.User
	if db.First(&user, id).Error != nil {
		content := gin.H{"error": "user with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	db.Delete(&user)

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
