package genreHandler

import "time"

type GenreCreateRequest struct {
	Name string `json:"name" form:"name"`
}

type GenreUpdateRequest struct {
	Name string `json:"name" form:"name"`
}

type GenreDataResponse struct {
	Id        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
