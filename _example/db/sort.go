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

func SortRecords(sorts string, db *gorm.DB) *gorm.DB {
	if sorts == "" {
		return db
	}

	for _, sort := range strings.Split(sorts, ",") {
		db = db.Order(convertPrefixToQuery(sort))
	}

	return db
}
