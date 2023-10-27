package comment

import (
	"mini_project/features/user"
	"time"
)

type Core struct {
	Id        string
	Comment   string
	GameId    string
	UserId    string
	User      user.Core
	CanEdit   bool
	CanDelete bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DataInterface interface {
	SelectAll(gameId string) ([]Core, error)
	SelectById(id string) (Core, error)
	Insert(role string, data Core) (Core, error)
	Update(role string, data Core) error
	Delete(role string, data Core) error
}

type UseCaseInterface interface {
	GetAll(gameId string) ([]Core, error)
	GetById(id string) (Core, error)
	Insert(role string, data Core) (Core, error)
	Update(role string, data Core) error
	Delete(role string, data Core) error
}
