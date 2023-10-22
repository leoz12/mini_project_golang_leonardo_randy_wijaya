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

func (repo *adminRepository) Insert(data admin.AdminCore) error {
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

func (repo *adminRepository) CheckByEmail(email string) (*admin.AdminCore, error) {
	var data Admin

	tx := repo.db.Where("email = ?", email).First(&data)

	if tx.Error == gorm.ErrRecordNotFound {
		return &admin.AdminCore{}, errors.New("invalid email or password")

	} else if tx.Error != nil {
		return &admin.AdminCore{}, tx.Error
	}

	return &admin.AdminCore{
		ID:        data.ID,
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
