package gameHandler

import (
	"mini_project/features/game"
	"time"
)

type CreateRequest struct {
	Name        string  `json:"name" form:"name"`
	Description string  `json:"description" form:"description"`
	Price       float32 `json:"price" form:"price"`
	Stock       int     `json:"stock" form:"stock"`
	Discount    float32 `json:"discount" form:"discount"`
	Genre       string  `json:"genre" form:"genre"`
	Publisher   string  `json:"publisher" form:"publisher"`
	ReleaseDate string  `json:"releaseDate" form:"releaseDate"`
}

type UpdateRequest struct {
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	Price       float32   `json:"price" form:"price"`
	Stock       int       `json:"stock" form:"stock"`
	Discount    float32   `json:"discount" form:"discount"`
	Genre       string    `json:"genre" form:"genre"`
	Publisher   string    `json:"publisher" form:"publisher"`
	ReleaseDate time.Time `json:"releaseDate" form:"releaseDate"`
}

type GameLiteResponse struct {
	Id          string
	Name        string
	Description string
	Price       float32
	Discount    float32
	Genre       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type GameResponse struct {
	Id          string
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

func GetGameLiteResponse(data game.GameCore) GameLiteResponse {
	return GameLiteResponse{
		Id:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Discount:    data.Discount,
		Genre:       data.Genre,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func GetGameResponse(data *game.GameCore) GameResponse {
	return GameResponse{
		Id:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Stock:       data.Stock,
		Discount:    data.Discount,
		Genre:       data.Genre,
		Publisher:   data.Publisher,
		ReleaseDate: data.ReleaseDate,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}
