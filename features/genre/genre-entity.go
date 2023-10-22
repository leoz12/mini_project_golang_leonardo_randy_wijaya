package genre

import (
	"time"
)

type GenreCore struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DataInterface interface {
	SelectAll() ([]GenreCore, error)
	Insert(data GenreCore) (*GenreCore, error)
	Update(id string, data GenreCore) error
	Delete(id string) error
}

type UseCaseInterface interface {
	GetAll() ([]GenreCore, error)
	Insert(data GenreCore) (*GenreCore, error)
	Update(id string, data GenreCore) error
	Delete(id string) error
}
