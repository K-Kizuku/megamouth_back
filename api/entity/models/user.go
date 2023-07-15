package models

import (
	"time"

	_ "gorm.io/gorm"
)

type User struct {
	ID             string          `gorm:"primaryKey" json:"id"`
	Name           string          `gorm:"not null" json:"name"`
	Follows        []*User         `gorm:"many2many:followings;foreignKey:ID;association_foreignKey:ID;joinForeignKey:follow_id;JoinReferences:follower_id" json:"follows"`
	Followers      []*User         `gorm:"many2many:followings;foreignKey:ID;association_foreignKey:ID;joinForeignKey:follower_id;JoinReferences:follow_id" json:"followers"`
	CreatedAt      time.Time       `gorm:"not null" json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	DeletedAt      time.Time       `json:"deleted_at"`
	PasswordHash   string          `gorm:"not null" json:"-"`
	RefreshToken   string          `gorm:"not null" json:"refresh_token"`
	Authentication *Authentication `gorm:"not null;foreignkey:ID"`
}

type Authentication struct {
	ID           string `gorm:"primaryKey" json:"id"`
	PasswordHash string `gorm:"not null" json:"-"`
	RefreshToken string `gorm:"not null" json:"refresh_token"`
}
