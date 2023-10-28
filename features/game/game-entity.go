package game

import (
	"mini_project/features/genre"
	"time"
)

type Core struct {
	Id          string
	Name        string
	Description string
	Price       float32
	Stock       int
	Discount    float32
	Genres      []genre.Core
	Publisher   string
	ImageUrl    string
	Platform    string
	ReleaseDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type GameParams struct {
	Genres string
	Search string
}

type DataInterface interface {
	SelectAll(GameParams) ([]Core, error)
	SelectById(id string) (Core, error)
	Insert(data Core) (Core, error)
	Update(id string, data Core) error
	Delete(id string) error
}

type UseCaseInterface interface {
	GetAll(GameParams) ([]Core, error)
	GetById(id string) (Core, error)
	Insert(data Core) (Core, error)
	Update(id string, data Core) error
	Delete(id string) error
}
