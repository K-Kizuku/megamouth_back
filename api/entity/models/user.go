package models

type User struct {
	ID string `gorm:"primaryKey" json:"id"`

	Name         string  `gorm:"not null" json:"name"`
	Email        string  `gorm:"unique;not null" json:"email"`
	PasswordHash string  `gorm:"not null" json:"-"`
	Follows      []*User `gorm:"many2many:user_followers;foreignKey:id;association_foreignKey:id;joinForeignKey:follow_id;JoinReferences:follower_id" json:"follows"`
	Followers    []*User `gorm:"many2many:user_followers;foreignKey:id;association_foreignKey:id;joinForeignKey:follower_id;JoinReferences:follow_id" json:"followers"`
}
