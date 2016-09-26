package db

import (
	"strings"

	"github.com/jinzhu/gorm"
)

func convertPrefixToQuery(sort string) string {
	if strings.HasPrefix(sort, "-") {
		return strings.TrimLeft(sort, "-") + " desc"
	} else {
		return strings.TrimLeft(sort, " ") + " asc"
	}
}

func (self *Parameter) SortRecords(db *gorm.DB) *gorm.DB {
	if self.Sort == "" {
		return db
	}

	for _, sort := range strings.Split(self.Sort, ",") {
		db = db.Order(convertPrefixToQuery(sort))
	}

	return db
}
