package Models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `gorm:"unique" json:"email"`
	Password  string
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
