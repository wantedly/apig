package models

type Email struct {
	ID      uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id" form:"id"`
	Address string `json:"address" form:"address"`
	UserID  uint   `json:"user_id" form:"user_id"`
	User    *User  `json:"user" form:"user"`
}
