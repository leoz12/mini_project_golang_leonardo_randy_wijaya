package genreHandler

import "time"

type CreateRequest struct {
	Name string `json:"name" form:"name"`
}

type UpdateRequest struct {
	Name string `json:"name" form:"name"`
}

type GenreDataResponse struct {
	Id        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
