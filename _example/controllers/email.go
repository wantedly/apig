package controllers

import (
	"encoding/json"
	"net/http"

	dbpkg "github.com/wantedly/apig/_example/db"
	"github.com/wantedly/apig/_example/helper"
	"github.com/wantedly/apig/_example/models"
	"github.com/wantedly/apig/_example/version"

	"github.com/gin-gonic/gin"
)

func GetEmails(c *gin.Context) {
	ver, err := version.NewVersion(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	parameter, err := dbpkg.NewParameter(c, models.Email{})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db, err = parameter.Paginate(db)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db = parameter.SetPreloads(db)
	db = parameter.SortRecords(db)
	db = parameter.FilterFields(db)
	emails := []models.Email{}
	fields := helper.ParseFields(c.DefaultQuery("fields", "*"))
	queryFields := helper.QueryFields(models.Email{}, fields)

	if err := db.Select(queryFields).Find(&emails).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	index := 0

	if len(emails) > 0 {
		index = int(emails[len(emails)-1].ID)
	}

	if err := parameter.SetHeaderLink(c, index); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if ver.Range("=>", "1.0.0") && ver.Range("<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	if _, ok := c.GetQuery("stream"); ok {
		enc := json.NewEncoder(c.Writer)
		c.Status(200)

		for _, email := range emails {
			fieldMap, err := helper.FieldToMap(email, fields)
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}

			if err := enc.Encode(fieldMap); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
		}
	} else {
		fieldMaps := []map[string]interface{}{}

		for _, email := range emails {
			fieldMap, err := helper.FieldToMap(email, fields)
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}

			fieldMaps = append(fieldMaps, fieldMap)
		}

		if _, ok := c.GetQuery("pretty"); ok {
			c.IndentedJSON(200, fieldMaps)
		} else {
			c.JSON(200, fieldMaps)
		}
	}
}

func GetEmail(c *gin.Context) {
	ver, err := version.NewVersion(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	parameter, err := dbpkg.NewParameter(c, models.Email{})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db = parameter.SetPreloads(db)
	email := models.Email{}
	id := c.Params.ByName("id")
	fields := helper.ParseFields(c.DefaultQuery("fields", "*"))
	queryFields := helper.QueryFields(models.Email{}, fields)

	if err := db.Select(queryFields).First(&email, id).Error; err != nil {
		content := gin.H{"error": "email with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}

	fieldMap, err := helper.FieldToMap(email, fields)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if ver.Range("=>", "1.0.0") && ver.Range("<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	if _, ok := c.GetQuery("pretty"); ok {
		c.IndentedJSON(200, fieldMap)
	} else {
		c.JSON(200, fieldMap)
	}
}

func CreateEmail(c *gin.Context) {
	ver, err := version.NewVersion(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	email := models.Email{}

	if err := c.Bind(&email); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&email).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if ver.Range("=>", "1.0.0") && ver.Range("<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(201, email)
}

func UpdateEmail(c *gin.Context) {
	ver, err := version.NewVersion(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	email := models.Email{}

	if db.First(&email, id).Error != nil {
		content := gin.H{"error": "email with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}

	if err := c.Bind(&email); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&email).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if ver.Range("=>", "1.0.0") && ver.Range("<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(200, email)
}

func DeleteEmail(c *gin.Context) {
	ver, err := version.NewVersion(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	email := models.Email{}

	if db.First(&email, id).Error != nil {
		content := gin.H{"error": "email with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}

	if err := db.Delete(&email).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if ver.Range("=>", "1.0.0") && ver.Range("<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
