package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           string    `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"not null" json:"name"`
	Email        string    `gorm:"unique;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Follows      []*User   `gorm:"many2many:user_followers;foreignKey:id;association_foreignKey:id;joinForeignKey:follow_id;JoinReferences:follower_id" json:"follows"`
	Followers    []*User   `gorm:"many2many:user_followers;foreignKey:id;association_foreignKey:id;joinForeignKey:follower_id;JoinReferences:follow_id" json:"followers"`
	Image        []*Image  `gorm:"foreignKey:UserID" json:"image"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

type Image struct {
	gorm.Model
	UserID string `json:"user_id"`
	URL    string `json:"url"`
}
