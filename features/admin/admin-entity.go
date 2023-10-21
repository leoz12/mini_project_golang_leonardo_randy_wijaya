package admin

import "time"

type AdminCore struct {
	ID        string
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
	Insert(data AdminCore) error
	CheckByEmail(email string) (*AdminCore, error)
}

type UseCaseInterface interface {
	Create(data AdminCore) error
	Login(data LoginCore) (string, error)
}
