package userHandler

import (
	"mini_project/features/user"
	"time"
)

type UserRegisterRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserResponse struct {
	Id        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CoreToResponse(data user.Core) UserResponse {
	return UserResponse{
		Id:        data.Id,
		Name:      data.Name,
		Email:     data.Email,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
