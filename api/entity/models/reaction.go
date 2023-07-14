package models

import (
	"gorm.io/gorm"
)

type Reaction struct {
	gorm.Model
	ReactionID uint   `gorm:"not null" json:"reaction_id"`
	User       *User  `gorm:"not null;foreignkey:UserID"`
	Post       *Post  `gorm:"not null;foreignkey:PostID"`
	UserID     string `gorm:"not null" json:"user_id"`
	PostID     string `gorm:"not null" json:"post_id"`
}
