package gameRepository

import (
	genreRepository "mini_project/features/genre/repository"
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
	Publisher   string
	Genres      []genreRepository.Genre `gorm:"many2many:game_genres;"`
	ImageUrl    string
	Platform    string
	ReleaseDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
