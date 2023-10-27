package transaction

import (
	userRepository "mini_project/features/user/repository"
	"time"
)

type Core struct {
	Id              string
	UserId          string
	User            userRepository.User
	GameId          string
	GameName        string
	GameDescription string
	Quantity        int
	Price           float32
	Discount        float32
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type DataInterface interface {
	SelectAll(userId string, role string) ([]Core, error)
	SelectById(id string) (Core, error)
	Insert(data Core) (Core, error)
}

type UseCaseInterface interface {
	GetAll(userId string, role string) ([]Core, error)
	GetById(id string) (Core, error)
	Insert(data Core) (Core, error)
}
