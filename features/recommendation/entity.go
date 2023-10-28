package recommendation

import "time"

type Core struct {
	Id        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DataInterface interface {
	SelectAll() ([]Core, error)
	SelectById(id string) (Core, error)
	Insert(data Core) (Core, error)
	Update(id string, data Core) error
	Delete(id string) error
}

type UseCaseInterface interface {
	RecommendGame(id string) (string, error)
	GetAll() ([]Core, error)
	GetById(id string) (Core, error)
	Insert(data Core) (Core, error)
	Update(id string, data Core) error
	Delete(id string) error
}
