package adminHandler

import (
	"mini_project/features/admin"
	"time"
)

type AdminRegisterRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type AdminLoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type AdminResponse struct {
	Id        string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CoreToResponse(data admin.Core) AdminResponse {
	return AdminResponse{
		Id:        data.Id,
		Email:     data.Email,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
