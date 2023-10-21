package repository

import (
	"mini_project/features/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (repo *userRepository) Insert(data user.UserCore) error {
	var input = User{
		ID:       uuid.New().String(),
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
	tx := repo.db.Create(&input)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *userRepository) CheckByEmail(email string) (*user.UserCore, error) {
	var data User

	tx := repo.db.Where("email = ?", email).First(&data)

	if tx.Error != nil {
		return &user.UserCore{}, tx.Error
	}

	return &user.UserCore{
		ID:        data.ID,
		Name:      data.Name,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}, nil
}

func UserDB(db *gorm.DB) user.DataInterface {
	return &userRepository{
		db: db,
	}
}
