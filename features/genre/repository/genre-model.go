package genreRepository

import (
	"time"

	"gorm.io/gorm"
)

type Genre struct {
	ID        string `gorm:"primarykey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
