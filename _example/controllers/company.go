package controllers

import (
	"net/http"
	"strings"

	dbpkg "github.com/wantedly/apig/_example/db"
	"github.com/wantedly/apig/_example/helper"
	"github.com/wantedly/apig/_example/models"
	"github.com/wantedly/apig/_example/version"

	"github.com/gin-gonic/gin"
)

func GetCompanies(c *gin.Context) {
	ver, err := version.New(c)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ids := c.DefaultQuery("ids", "")
	preloads := c.DefaultQuery("preloads", "")
	fields := helper.ParseFields(c.DefaultQuery("fields", "*"))
	queryFields := helper.QueryFields(models.Company{}, fields)

	pagination := dbpkg.Pagination{}
	db, err := pagination.Paginate(c)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db = dbpkg.SetPreloads(preloads, db)

	if ids != "" {
		db = db.Where("id IN (?)", strings.Split(ids, ","))
	}

	var companies []models.Company
	if err := db.Select(queryFields).Find(&companies).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	index := 0

	if len(companies) > 0 {
		index = int(companies[len(companies)-1].ID)
	}

	if err := pagination.SetHeaderLink(c, index); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	fieldMaps := []map[string]interface{}{}
	for _, company := range companies {
		fieldMap, err := helper.FieldToMap(company, fields)

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

func GetCompany(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := c.Params.ByName("id")
	preloads := c.DefaultQuery("preloads", "")
	fields := helper.ParseFields(c.DefaultQuery("fields", "*"))
	queryFields := helper.QueryFields(models.Company{}, fields)

	db := dbpkg.DBInstance(c)
	db = dbpkg.SetPreloads(preloads, db)

	var company models.Company
	if err := db.Select(queryFields).First(&company, id).Error; err != nil {
		content := gin.H{"error": "company with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	fieldMap, err := helper.FieldToMap(company, fields)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if _, ok := c.GetQuery("pretty"); ok {
		c.IndentedJSON(200, fieldMap)
	} else {
		c.JSON(200, fieldMap)
	}
}

func CreateCompany(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	var company models.Company

	if err := c.Bind(&company); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&company).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
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

	if err := c.Bind(&company); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&company).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

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

	if err := db.Delete(&company).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
