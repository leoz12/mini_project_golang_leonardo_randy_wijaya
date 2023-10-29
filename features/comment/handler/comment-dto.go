package commentHandler

import (
	"mini_project/features/comment"
	"time"
)

type CreateRequest struct {
	GameId  string `json:"gameId" form:"gameId" validate:"required"`
	Comment string `json:"comment" form:"comment" validate:"required"`
}

type UpdateRequest struct {
	Comment string `json:"comment" form:"comment" validate:"required"`
}

type CommentUser struct {
	Id   string
	Name string
}
type CommentResponse struct {
	Id        string
	Comment   string
	GameId    string
	User      CommentUser
	CanEdit   bool
	CanDelete bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CommentLiteResponse struct {
	Id        string
	Comment   string
	GameId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CoreToResponse(data comment.Core, userId string) CommentResponse {
	var canEdit, canDelete bool

	if data.UserId == userId {
		canEdit = true
		canDelete = true
	}
	return CommentResponse{
		Id:      data.Id,
		Comment: data.Comment,
		User: CommentUser{
			Id:   data.User.Id,
			Name: data.User.Name,
		},
		CanEdit:   canEdit,
		CanDelete: canDelete,
		GameId:    data.GameId,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func CoreToResponseLite(data comment.Core) CommentLiteResponse {
	return CommentLiteResponse{
		Id:        data.Id,
		Comment:   data.Comment,
		GameId:    data.GameId,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
