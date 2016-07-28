package controllers

import (
	"net/http"

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

	preloads := c.DefaultQuery("preloads", "")
	fields := helper.ParseFields(c.DefaultQuery("fields", "*"))

	pagination := dbpkg.Pagination{}
	db, err := pagination.Paginate(c)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db = dbpkg.SetPreloads(preloads, db)

	var jobs []models.Job
	if err := db.Select("*").Find(&jobs).Error; err != nil {
		c.JSON(500, gin.H{"error": "error occured"})
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

	var fieldMap []map[string]interface{}
	for key, _ := range jobs {
		fieldMap = append(fieldMap, helper.FieldToMap(jobs[key], fields))
	}
	c.JSON(200, fieldMap)
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

	db := dbpkg.DBInstance(c)
	db = dbpkg.SetPreloads(preloads, db)

	var job models.Job
	if err := db.Select("*").First(&job, id).Error; err != nil {
		content := gin.H{"error": "job with id#" + id + " not found"}
		c.JSON(404, content)
		return
	}

	if version.Range("1.0.0", "<=", ver) && version.Range(ver, "<", "2.0.0") {
		// conditional branch by version.
		// 1.0.0 <= this version < 2.0.0 !!
	}

	fieldMap := helper.FieldToMap(job, fields)
	c.JSON(200, fieldMap)
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
		content := gin.H{"error": "error occured"}
		c.JSON(500, content)
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
