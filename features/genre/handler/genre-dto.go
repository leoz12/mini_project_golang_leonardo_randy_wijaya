package genreHandler

import (
	"mini_project/features/genre"
	"time"
)

type CreateRequest struct {
	Name string `json:"name" form:"name"`
}

type UpdateRequest struct {
	Name string `json:"name" form:"name"`
}

type GenreResponse struct {
	Id        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CoreToResponse(data genre.Core) GenreResponse {
	return GenreResponse{
		Id:        data.Id,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
