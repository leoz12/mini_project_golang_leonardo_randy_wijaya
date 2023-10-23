package game

import (
	"time"
)

type GameCore struct {
	ID          string
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
}

type DataInterface interface {
	SelectAll() ([]GameCore, error)
	SelectById(id string) (*GameCore, error)
	Insert(data GameCore) (*GameCore, error)
	Update(id string, data GameCore) error
	Delete(id string) error
}

type UseCaseInterface interface {
	GetAll() ([]GameCore, error)
	GetById(id string) (*GameCore, error)
	Insert(data GameCore) (*GameCore, error)
	Update(id string, data GameCore) error
	Delete(id string) error
}
