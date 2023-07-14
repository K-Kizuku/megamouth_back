package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	User    *User  `gorm:"not null;foreignkey:Author"`
	Author  string `gorm:"not null" json:"author"`
	Content string `gorm:"not null" json:"content"`
	reply   uint   `gorm:"foreignkey:ID"`
}
