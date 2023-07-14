package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	User  *User  `gorm:"not null;foreignkey:Owner"`
	Owner string `json:"owner"`
	URL   string `gorm:"not null" json:"url"`
}
