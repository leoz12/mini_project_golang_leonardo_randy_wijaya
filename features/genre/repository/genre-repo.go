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

// Insert implements genre.DataInterface.
func (repo *genreRepository) Insert(data genre.GenreCore) (*genre.GenreCore, error) {
	var input = Genre{
		ID:   uuid.New().String(),
		Name: data.Name,
	}
	tx := repo.db.Create(&input)

	if tx.Error != nil {
		return &genre.GenreCore{
			ID:        input.ID,
			Name:      input.Name,
			CreatedAt: input.CreatedAt,
			UpdatedAt: input.UpdatedAt,
		}, tx.Error
	}
	return &genre.GenreCore{
		ID:        input.ID,
		Name:      input.Name,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}, nil
}

// SelectAll implements genre.DataInterface.
func (repo *genreRepository) SelectAll() ([]genre.GenreCore, error) {

	var genres []Genre

	tx := repo.db.Find(&genres)

	var genresCore []genre.GenreCore

	if tx.Error != nil {
		return genresCore, tx.Error
	}
	for _, val := range genres {
		genresCore = append(genresCore, genre.GenreCore{
			ID:        val.ID,
			Name:      val.Name,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}
	return genresCore, nil
}

// Update implements genre.DataInterface.
func (repo *genreRepository) Update(id string, data genre.GenreCore) error {
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

func New(db *gorm.DB) genre.DataInterface {
	return &genreRepository{
		db: db,
	}
}
