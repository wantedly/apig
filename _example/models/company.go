package models

import "database/sql"

type Company struct {
	ID   uint           `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name string         `json:"name"`
	URL  sql.NullString `json:"url"`
	Jobs []*Job         `json:"jobs"`
}
