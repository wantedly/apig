package models

type User struct {
	ID      uint     `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name    string   `json:"name"`
	Profile *Profile `json:"profile"`
	Jobs    []*Job   `json:"jobs"`
	Emails  []*Email `json:"emails"`
}
