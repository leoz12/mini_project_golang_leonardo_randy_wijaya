package wishlist

import (
	gameRepository "mini_project/features/game/repository"
	"time"
)

type WishlistCore struct {
	Id        string
	GameId    string
	Game      gameRepository.Game
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DataInterface interface {
	SelectAll(userId string) ([]WishlistCore, error)
	Insert(data WishlistCore) (WishlistCore, error)
	Delete(id string) error
}

type UseCaseInterface interface {
	GetAll(userId string) ([]WishlistCore, error)
	Insert(data WishlistCore) (WishlistCore, error)
	Delete(id string) error
}
