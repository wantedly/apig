package models

import "time"

type Profile struct {
	ID       uint      `gorm:"primary_key;AUTO_INCREMENT" json:"id" form:"id"`
	UserID   uint      `json:"user_id" form:"user_id"`
	User     *User     `json:"user" form:"user"`
	Birthday time.Time `json:"birthday" form:"birthday"`
	Engaged  bool      `json:"engaged" form:"engaged"`
}
