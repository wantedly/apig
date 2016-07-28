package models

import "time"

type Profile struct {
	ID       uint      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UserID   uint      `json:"user_id"`
	User     *User     `json:"user"`
	Birthday time.Time `json:"birthday"`
	Engaged  bool      `json:"engaged"`
}
