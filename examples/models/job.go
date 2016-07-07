package model

import (
	"time"
)

type Job struct {
	ID          uint       `gorm:"primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
