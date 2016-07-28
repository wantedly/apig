package models

type Job struct {
	ID        uint  `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UserID    uint  `json:"user_id"`
	User      *User `json:"user"`
	CompanyID uint  `json:"company_id"`
	RoleCD    uint  `json:"role_cd"`
}
