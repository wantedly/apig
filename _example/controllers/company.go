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

func setCompanyPreload(fields []string, db *gorm.DB) ([]string, *gorm.DB) {
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

		}
	}
	return sel, db
}

func GetCompanies(c *gin.Context) {
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
	sel, db := setCompanyPreload(fields, db)
	var companies []models.Company
	err = db.Select(sel).Find(&companies).Error

	if err != nil {
		c.JSON(500, gin.H{"error": "error occured"})
		return
	}

	// paging
	var index int
	if len(companies) < 1 {
		index = 0
	} else {
		index = int(companies[len(companies)-1].ID)
	}
	pagination.SetHeaderLink(c, index)

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	var fieldMap []map[string]interface{}
	for key, _ := range companies {
		fieldMap = append(fieldMap, helper.FieldToMap(companies[key], fields, nestFields))
	}
	c.JSON(200, fieldMap)
}

func GetCompany(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	fields, nestFields := helper.ParseFields(c.DefaultQuery("fields", "*"))
	sel, db := setCompanyPreload(fields, db)
	var company models.Company

	if db.Select(sel).First(&company, id).Error != nil {
		content := gin.H{"error": "company with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	fieldMap := helper.FieldToMap(company, fields, nestFields)
	c.JSON(200, fieldMap)
}

func CreateCompany(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	var company models.Company
	c.Bind(&company)
	if db.Create(&company).Error != nil {
		content := gin.H{"error": "error occured"}
		c.JSON(500, content)
		return
	}

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(201, company)
}

func UpdateCompany(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	var company models.Company
	if db.First(&company, id).Error != nil {
		content := gin.H{"error": "company with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.Bind(&company)
	db.Save(&company)

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(200, company)
}

func DeleteCompany(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	var company models.Company
	if db.First(&company, id).Error != nil {
		content := gin.H{"error": "company with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	db.Delete(&company)

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
