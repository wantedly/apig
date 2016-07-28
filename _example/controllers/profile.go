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

func setProfilePreload(fields []string, db *gorm.DB) ([]string, *gorm.DB) {
	sel := make([]string, len(fields))
	copy(sel, fields)
	offset := 0
	for key, val := range fields {
		switch val {

		case "user":
			db = db.Preload("User")
			sel = append(sel[:(key-offset)], sel[(key+1-offset):]...)
			offset += 1
			idflag := true
			for _, v := range sel {
				if v == "user_id" {
					idflag = false
					break
				}
			}
			if idflag {
				sel = append(sel, "user_id")
			}

		}
	}
	return sel, db
}

func GetProfiles(c *gin.Context) {
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
	sel, db := setProfilePreload(fields, db)
	var profiles []models.Profile
	err = db.Select(sel).Find(&profiles).Error

	if err != nil {
		c.JSON(500, gin.H{"error": "error occured"})
		return
	}

	// paging
	var index int
	if len(profiles) < 1 {
		index = 0
	} else {
		index = int(profiles[len(profiles)-1].ID)
	}
	pagination.SetHeaderLink(c, index)

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	var fieldMap []map[string]interface{}
	for key, _ := range profiles {
		fieldMap = append(fieldMap, helper.FieldToMap(profiles[key], fields, nestFields))
	}
	c.JSON(200, fieldMap)
}

func GetProfile(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	fields, nestFields := helper.ParseFields(c.DefaultQuery("fields", "*"))
	sel, db := setProfilePreload(fields, db)
	var profile models.Profile

	if db.Select(sel).First(&profile, id).Error != nil {
		content := gin.H{"error": "profile with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	fieldMap := helper.FieldToMap(profile, fields, nestFields)
	c.JSON(200, fieldMap)
}

func CreateProfile(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	var profile models.Profile
	c.Bind(&profile)
	if db.Create(&profile).Error != nil {
		content := gin.H{"error": "error occured"}
		c.JSON(500, content)
		return
	}

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(201, profile)
}

func UpdateProfile(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	var profile models.Profile
	if db.First(&profile, id).Error != nil {
		content := gin.H{"error": "profile with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.Bind(&profile)
	db.Save(&profile)

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(200, profile)
}

func DeleteProfile(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	var profile models.Profile
	if db.First(&profile, id).Error != nil {
		content := gin.H{"error": "profile with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	db.Delete(&profile)

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
