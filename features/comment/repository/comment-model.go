package commentRepository

import (
	"mini_project/features/comment"
	"mini_project/features/user"
	userRepository "mini_project/features/user/repository"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        string `gorm:"primarykey"`
	Comment   string
	GameID    string              `gorm:"size:191"`
	UserID    string              `gorm:"size:191"`
	User      userRepository.User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func ModelToCore(data Comment) comment.Core {
	return comment.Core{
		Id:      data.ID,
		Comment: data.Comment,
		UserId:  data.UserID,
		User: user.Core{
			Id:    data.User.ID,
			Name:  data.User.Name,
			Email: data.User.Email,
		},
		GameId:    data.GameID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
