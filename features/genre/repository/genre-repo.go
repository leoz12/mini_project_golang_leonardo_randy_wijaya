package genreRepository

import (
	"errors"
	"mini_project/features/genre"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type genreRepository struct {
	db *gorm.DB
}

// SelectAll implements genre.DataInterface.
func (repo *genreRepository) SelectAll() ([]genre.Core, error) {

	var genres []Genre

	tx := repo.db.Find(&genres)

	var genresCore []genre.Core

	if tx.Error != nil {
		return genresCore, tx.Error
	}
	for _, val := range genres {
		genresCore = append(genresCore, genre.Core{
			Id:        val.ID,
			Name:      val.Name,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}
	return genresCore, nil
}

// SelectById implements genre.DataInterface.
func (repo *genreRepository) SelectById(id string) (genre.Core, error) {
	var data Genre

	tx := repo.db.Where("id = ?", id).First(&data)

	if tx.RowsAffected == 0 {
		return genre.Core{}, errors.New("invalid id")
	}

	if tx.Error != nil {
		return genre.Core{}, tx.Error
	}
	return genre.Core{
		Id:        data.ID,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}, nil
}

// Insert implements genre.DataInterface.
func (repo *genreRepository) Insert(data genre.Core) (genre.Core, error) {
	var input = Genre{
		ID:   uuid.New().String(),
		Name: data.Name,
	}
	tx := repo.db.Create(&input)

	if tx.Error != nil {
		return genre.Core{}, tx.Error
	}
	return genre.Core{
		Id:        input.ID,
		Name:      input.Name,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}, nil
}

// Update implements genre.DataInterface.
func (repo *genreRepository) Update(id string, data genre.Core) error {
	tx := repo.db.Model(&Genre{
		ID: id,
	}).Update("name", data.Name)

	if tx.RowsAffected == 0 {
		return errors.New("invalid id")
	}
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements genre.DataInterface.
func (repo *genreRepository) Delete(id string) error {

	tx := repo.db.Where("id = ?", id).Delete(&Genre{})

	if tx.RowsAffected == 0 {
		return errors.New("invalid id")
	}
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func New(db *gorm.DB) genre.DataInterface {
	return &genreRepository{
		db: db,
	}
}
