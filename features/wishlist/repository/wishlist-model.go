package wishlistRepository

import (
	gameRepository "mini_project/features/game/repository"
	"mini_project/features/wishlist"
	"time"

	"gorm.io/gorm"
)

type Wishlist struct {
	ID        string              `gorm:"primarykey"`
	GameID    string              `gorm:"size:191"`
	Game      gameRepository.Game `gorm:"foreignKey:GameID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    string              `gorm:"size:191"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func ModelToCore(data Wishlist) wishlist.WishlistCore {
	return wishlist.WishlistCore{
		Id:        data.ID,
		UserId:    data.UserID,
		GameId:    data.GameID,
		Game:      data.Game,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
