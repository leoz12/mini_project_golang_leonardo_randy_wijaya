package user

import (
	"time"

	"gorm.io/gorm"
)

type UserCore struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  gorm.DeletedAt
}

type LoginCore struct {
	Email    string
	Password string
}

type DataInterface interface {
	Insert(data UserCore) error
	CheckByEmail(email string) (*UserCore, error)
}

type UseCaseInterface interface {
	Create(data UserCore) error
	Login(data LoginCore) (string, error)
}
