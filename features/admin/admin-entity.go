package admin

import "time"

type Core struct {
	Id        string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DataInterface interface {
	SelectAll() ([]Core, error)
	Insert(data Core) error
	CheckByEmail(email string) (Core, error)
}

type UseCaseInterface interface {
	GetAll() ([]Core, error)
	Create(data Core) error
	Login(data Core) (string, error)
}
