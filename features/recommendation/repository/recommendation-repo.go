package recommendationRepository

import (
	"errors"
	"mini_project/features/recommendation"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type recommendationRepository struct {
	db *gorm.DB
}

// SelectAll implements recommendation.DataInterface.
func (repo *recommendationRepository) SelectAll() ([]recommendation.Core, error) {
	var genres []Recommendation

	tx := repo.db.Find(&genres)

	var genresCore []recommendation.Core

	if tx.Error != nil {
		return genresCore, tx.Error
	}
	for _, val := range genres {
		genresCore = append(genresCore, recommendation.Core{
			Id:        val.ID,
			Name:      val.Name,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}
	return genresCore, nil
}

// SelectById implements recommendation.DataInterface.
func (repo *recommendationRepository) SelectById(id string) (recommendation.Core, error) {
	var data Recommendation

	tx := repo.db.Where("id = ?", id).First(&data)

	if tx.RowsAffected == 0 {
		return recommendation.Core{}, errors.New("invalid id")
	}

	if tx.Error != nil {
		return recommendation.Core{}, tx.Error
	}
	return recommendation.Core{
		Id:        data.ID,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}, nil
}

// Insert implements recommendation.DataInterface.
func (repo *recommendationRepository) Insert(data recommendation.Core) (recommendation.Core, error) {
	var input = Recommendation{
		ID:   uuid.New().String(),
		Name: data.Name,
	}
	tx := repo.db.Create(&input)

	if tx.Error != nil {
		return recommendation.Core{}, tx.Error
	}
	return recommendation.Core{
		Id:        input.ID,
		Name:      input.Name,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}, nil
}

// Update implements recommendation.DataInterface.
func (repo *recommendationRepository) Update(id string, data recommendation.Core) error {
	tx := repo.db.Model(&Recommendation{
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

// Delete implements recommendation.DataInterface.
func (repo *recommendationRepository) Delete(id string) error {

	tx := repo.db.Where("id = ?", id).Delete(&Recommendation{})

	if tx.RowsAffected == 0 {
		return errors.New("invalid id")
	}
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func New(db *gorm.DB) recommendation.DataInterface {
	return &recommendationRepository{
		db: db,
	}
}
