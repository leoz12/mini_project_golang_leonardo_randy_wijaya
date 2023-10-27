package adminRepository

import (
	"errors"
	"mini_project/features/admin"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

func (repo *adminRepository) SelectAll() ([]admin.Core, error) {
	var admins []Admin
	var adminsCore []admin.Core
	tx := repo.db.Find(&admins)

	if tx.Error != nil {
		return adminsCore, tx.Error
	}
	for _, val := range admins {
		adminsCore = append(adminsCore, admin.Core{
			Id:        val.ID,
			Email:     val.Email,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}
	return adminsCore, nil

}

func (repo *adminRepository) Insert(data admin.Core) error {
	var input = Admin{
		ID:       uuid.New().String(),
		Email:    data.Email,
		Password: data.Password,
	}
	tx := repo.db.Create(&input)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *adminRepository) CheckByEmail(email string) (admin.Core, error) {
	var data Admin

	tx := repo.db.Where("email = ?", email).First(&data)

	if tx.Error == gorm.ErrRecordNotFound {
		return admin.Core{}, errors.New("invalid email or password")

	} else if tx.Error != nil {
		return admin.Core{}, tx.Error
	}

	return admin.Core{
		Id:        data.ID,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}, nil
}

func New(db *gorm.DB) admin.DataInterface {
	return &adminRepository{
		db: db,
	}
}
