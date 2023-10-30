package transactionRepository

import (
	userRepository "mini_project/features/user/repository"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID              string              `gorm:"primarykey"`
	UserId          string              `gorm:"size:191"`
	User            userRepository.User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	GameId          string
	GameName        string
	GameDescription string
	Quantity        int
	Price           float32
	Discount        float32
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
