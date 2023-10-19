package user

import "time"

type UserCore struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
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
