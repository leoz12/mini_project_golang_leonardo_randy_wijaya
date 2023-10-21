package repository

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        string `gorm:"primarykey"`
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
