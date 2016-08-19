package models

type Job struct {
	ID        uint  `gorm:"primary_key;AUTO_INCREMENT" json:"id" form:"id"`
	UserID    uint  `json:"user_id" form:"user_id"`
	User      *User `json:"user" form:"user"`
	CompanyID uint  `json:"company_id" form:"company_id"`
	RoleCD    uint  `json:"role_cd" form:"role_cd"`
}
