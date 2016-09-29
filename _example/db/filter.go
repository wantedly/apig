package db

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func filterToMap(c *gin.Context, model interface{}) map[string]string {
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

	return filters
}

func (self *Parameter) FilterFields(db *gorm.DB) *gorm.DB {
	for k, v := range self.Filters {
		if v != "" {
			db = db.Where(fmt.Sprintf("%s IN (?)", k), strings.Split(v, ","))
		}
	}

	return db
}

func (self *Parameter) GetRawFilterQuery() string {
	var s string

	for k, v := range self.Filters {
		if v != "" {
			s += "&q[" + k + "]=" + v
		}
	}

	return s
}
