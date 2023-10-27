package adminHandler

import (
	"mini_project/features/admin"
	"time"
)

type AdminRegisterRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type AdminLoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
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
