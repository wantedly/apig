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

func GetJobs(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ids := c.DefaultQuery("ids", "")
	preloads := c.DefaultQuery("preloads", "")
	fields := helper.ParseFields(c.DefaultQuery("fields", "*"))
	queryFields := helper.QueryFields(models.Job{}, fields)

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

	var jobs []models.Job
	if err := db.Select(queryFields).Find(&jobs).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// paging
	var index int
	if len(jobs) < 1 {
		index = 0
	} else {
		index = int(jobs[len(jobs)-1].ID)
	}
	pagination.SetHeaderLink(c, index)

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	fieldMap := []map[string]interface{}{}
	for _, job := range jobs {
		fieldMap = append(fieldMap, helper.FieldToMap(job, fields))
	}

	_, ok := c.GetQuery("pretty")
	if ok {
		c.IndentedJSON(200, fieldMap)
	} else {
		c.JSON(200, fieldMap)
	}
}

func GetJob(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := c.Params.ByName("id")
	preloads := c.DefaultQuery("preloads", "")
	fields := helper.ParseFields(c.DefaultQuery("fields", "*"))
	queryFields := helper.QueryFields(models.Job{}, fields)

	db := dbpkg.DBInstance(c)
	db = dbpkg.SetPreloads(preloads, db)

	var job models.Job
	if err := db.Select(queryFields).First(&job, id).Error; err != nil {
		content := gin.H{"error": "job with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	fieldMap := helper.FieldToMap(job, fields)

	_, ok := c.GetQuery("pretty")
	if ok {
		c.IndentedJSON(200, fieldMap)
	} else {
		c.JSON(200, fieldMap)
	}
}

func CreateJob(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	var job models.Job
	c.Bind(&job)
	if db.Create(&job).Error != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(201, job)
}

func UpdateJob(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	var job models.Job
	if db.First(&job, id).Error != nil {
		content := gin.H{"error": "job with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	c.Bind(&job)
	db.Save(&job)

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.JSON(200, job)
}

func DeleteJob(c *gin.Context) {
	ver, err := version.New(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbpkg.DBInstance(c)
	id := c.Params.ByName("id")
	var job models.Job
	if db.First(&job, id).Error != nil {
		content := gin.H{"error": "job with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}
	db.Delete(&job)

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
