package wishlistHandler

import (
	gameHandler "mini_project/features/game/handler"
	"mini_project/features/wishlist"
	"time"
)

type CreateRequest struct {
	GameId string `json:"gameId" form:"gameId"`
}

type WishlistResponse struct {
	Id        string
	Game      gameHandler.GameResponse
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetWishlistReponse(data wishlist.WishlistCore) WishlistResponse {
	return WishlistResponse{
		Id: data.Id,
		Game: gameHandler.GameResponse{
			Id:          data.Game.ID,
			Name:        data.Game.Name,
			Description: data.Game.Description,
			Price:       data.Game.Price,
			Stock:       data.Game.Stock,
			Discount:    data.Game.Discount,
			Genre:       data.Game.Genre,
			Publisher:   data.Game.Publisher,
			ReleaseDate: data.Game.ReleaseDate,
			CreatedAt:   data.Game.CreatedAt,
			UpdatedAt:   data.Game.UpdatedAt,
		},
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
