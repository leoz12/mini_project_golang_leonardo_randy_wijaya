package userRepository

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primarykey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
