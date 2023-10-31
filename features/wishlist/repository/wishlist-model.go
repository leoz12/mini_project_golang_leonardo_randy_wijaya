package wishlistRepository

import (
	"mini_project/features/game"
	gameRepository "mini_project/features/game/repository"
	"mini_project/features/genre"
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

func ModelToCore(data Wishlist) wishlist.Core {
	var genres []genre.Core

	for _, val := range data.Game.Genres {
		genres = append(genres, genre.Core{
			Id:   val.ID,
			Name: val.Name,
		})
	}
	return wishlist.Core{
		Id:     data.ID,
		UserId: data.UserID,
		GameId: data.GameID,
		Game: game.Core{
			Id:          data.Game.ID,
			Name:        data.Game.Name,
			Description: data.Game.Description,
			Price:       data.Game.Price,
			Stock:       data.Game.Stock,
			Discount:    data.Game.Discount,
			Genres:      genres,
			ImageUrl:    data.Game.ImageUrl,
			Publisher:   data.Game.Publisher,
			CreatedAt:   data.Game.CreatedAt,
			UpdatedAt:   data.Game.UpdatedAt,
		},
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
