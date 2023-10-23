package gameRepository

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	ID          string `gorm:"primarykey"`
	Name        string
	Description string
	Price       float32
	Stock       int
	Discount    float32
	Genre       string
	Publisher   string
	ReleaseDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
