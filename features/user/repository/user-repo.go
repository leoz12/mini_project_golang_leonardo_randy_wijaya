package userRepository

import (
	"mini_project/features/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (repo *userRepository) SelectAll() ([]user.Core, error) {
	var users []User
	var usersCore []user.Core
	tx := repo.db.Find(&users)

	if tx.Error != nil {
		return usersCore, tx.Error
	}
	for _, val := range users {
		usersCore = append(usersCore, user.Core{
			Id:        val.ID,
			Name:      val.Name,
			Email:     val.Email,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}
	return usersCore, nil

}

func (repo *userRepository) Insert(data user.Core) error {
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

func (repo *userRepository) CheckByEmail(email string) (user.Core, error) {
	var data User

	tx := repo.db.Where("email = ?", email).First(&data)

	if tx.Error != nil {
		return user.Core{}, tx.Error
	}

	return user.Core{
		Id:        data.ID,
		Name:      data.Name,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}, nil
}

func New(db *gorm.DB) user.DataInterface {
	return &userRepository{
		db: db,
	}
}
