package gameHandler

import (
	"mini_project/features/game"
	"time"
)

type CreateRequest struct {
	Name        string   `json:"name" form:"name" validate:"required"`
	Description string   `json:"description" form:"description" validate:"required"`
	Price       float32  `json:"price" form:"price" validate:"required,numeric"`
	Stock       int      `json:"stock" form:"stock" validate:"required,numeric"`
	Discount    float32  `json:"discount" form:"discount" validate:"required,numeric,min=0,max=100"`
	Genres      []string `json:"genres" form:"genres" validate:"required,dive,required"`
	Publisher   string   `json:"publisher" form:"publisher" validate:"required"`
	ReleaseDate string   `json:"releaseDate" form:"releaseDate" validate:"required"`
	ImageUrl    string   `json:"ImageUrl" form:"ImageUrl"`
	Platform    string   `json:"Platform" form:"Platform" validate:"required"`
}

type UpdateRequest struct {
	Name        string   `json:"name" form:"name" validate:"required"`
	Description string   `json:"description" form:"description" validate:"required"`
	Price       float32  `json:"price" form:"price" validate:"required,numeric"`
	Stock       int      `json:"stock" form:"stock" validate:"required,numeric"`
	Discount    float32  `json:"discount" form:"discount" validate:"required,numeric,min=0,max=100"`
	Genres      []string `json:"genres" form:"genres" validate:"required,dive,required"`
	Publisher   string   `json:"publisher" form:"publisher" validate:"required"`
	ReleaseDate string   `json:"releaseDate" form:"releaseDate" validate:"required"`
	ImageUrl    string   `json:"ImageUrl" form:"ImageUrl"`
	Platform    string   `json:"Platform" form:"Platform" validate:"required"`
}

type GameGenre struct {
	Id   string
	Name string
}
type GameLiteResponse struct {
	Id          string
	Name        string
	Description string
	Price       float32
	Stock       int
	Discount    float32
	ImageUrl    string
	Platform    string
	Genres      []GameGenre
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
	ImageUrl    string
	Genres      []GameGenre
	Publisher   string
	Platform    string
	CanComment  bool
	ReleaseDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateGameResponse struct {
	Id          string
	Name        string
	Description string
	Price       float32
	Stock       int
	Discount    float32
	Publisher   string
	ImageUrl    string
	Platform    string
	ReleaseDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func CoreToLiteReponse(data game.Core) GameLiteResponse {
	var gameGenres []GameGenre

	for _, val := range data.Genres {
		gameGenres = append(gameGenres, GameGenre{
			Id:   val.Id,
			Name: val.Name,
		})
	}

	return GameLiteResponse{
		Id:          data.Id,
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Stock:       data.Stock,
		Discount:    data.Discount,
		Genres:      gameGenres,
		ImageUrl:    data.ImageUrl,
		Platform:    data.Platform,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func CoreToReponse(data game.Core) GameResponse {
	var gameGenres []GameGenre

	for _, val := range data.Genres {
		gameGenres = append(gameGenres, GameGenre{
			Id:   val.Id,
			Name: val.Name,
		})
	}

	return GameResponse{
		Id:          data.Id,
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Stock:       data.Stock,
		Discount:    data.Discount,
		Genres:      gameGenres,
		ImageUrl:    data.ImageUrl,
		Publisher:   data.Publisher,
		Platform:    data.Platform,
		CanComment:  data.CanComment,
		ReleaseDate: data.ReleaseDate,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func CoreToCreateReponse(data game.Core) CreateGameResponse {
	return CreateGameResponse{
		Id:          data.Id,
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Stock:       data.Stock,
		Discount:    data.Discount,
		Publisher:   data.Publisher,
		ImageUrl:    data.ImageUrl,
		Platform:    data.Platform,
		ReleaseDate: data.ReleaseDate,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}
