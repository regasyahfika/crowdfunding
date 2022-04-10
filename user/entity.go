package user

import "time"

type User struct {
	ID             int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name           string `json:"name"`
	Occupation     string `json:"occupation"`
	Email          string `gorm:"type:varchar(100)" json:"email"`
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
