package models

type User struct {
	ID      uint     `gorm:"primary_key;AUTO_INCREMENT" json:"id" form:"id"`
	Name    string   `json:"name" form:"name"`
	Profile *Profile `json:"profile" form:"profile"`
	Jobs    []*Job   `json:"jobs" form:"jobs"`
	Emails  []*Email `json:"emails" form:"emails"`
}
