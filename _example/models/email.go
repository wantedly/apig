package models

type Email struct {
	ID      uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Address string `json:"address"`
	UserID  uint   `json:"user_id"`
	User    *User  `json:"user"`
}
