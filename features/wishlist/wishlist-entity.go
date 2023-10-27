package wishlist

import (
	"mini_project/features/game"
	"time"
)

type Core struct {
	Id        string
	GameId    string
	Game      game.Core
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DataInterface interface {
	SelectAll(userId string) ([]Core, error)
	SelectById(id string) (Core, error)
	Insert(data Core) (Core, error)
	Delete(id string) error
}

type UseCaseInterface interface {
	GetAll(userId string) ([]Core, error)
	Insert(data Core) (Core, error)
	Delete(id string, userId string) error
}
