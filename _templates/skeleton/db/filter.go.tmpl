package db

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func FilterFields(c *gin.Context, model interface{}, db *gorm.DB) *gorm.DB {
	var jsonTag, jsonKey string
	filters := make(map[string]string)
	ts := reflect.TypeOf(model)

	for i := 0; i < ts.NumField(); i++ {
		f := ts.Field(i)
		jsonKey = f.Name

		if jsonTag = f.Tag.Get("json"); jsonTag != "" {
			jsonKey = strings.Split(jsonTag, ",")[0]
		}

		filters[jsonKey] = c.Query("q[" + jsonKey + "]")
	}

	for k, v := range filters {
		if v != "" {
			db = db.Where(fmt.Sprintf("%s IN (?)", k), strings.Split(v, ","))
		}
	}

	return db
}
