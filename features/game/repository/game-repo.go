package gameRepository

import (
	"errors"
	"mini_project/features/game"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type gameRepository struct {
	db *gorm.DB
}

// SelectAll implements game.DataInterface.
func (repo *gameRepository) SelectAll() ([]game.GameCore, error) {
	var games []Game

	tx := repo.db.Find(&games)

	var GamesCore []game.GameCore

	if tx.Error != nil {
		return GamesCore, tx.Error
	}
	for _, val := range games {
		GamesCore = append(GamesCore, game.GameCore{
			ID:          val.ID,
			Name:        val.Name,
			Description: val.Description,
			Price:       val.Price,
			Stock:       val.Stock,
			Discount:    val.Discount,
			Genre:       val.Genre,
			Publisher:   val.Publisher,
			ReleaseDate: val.ReleaseDate,
			CreatedAt:   val.CreatedAt,
			UpdatedAt:   val.UpdatedAt,
		})
	}
	return GamesCore, nil
}

// SelectById implements game.DataInterface.
func (repo *gameRepository) SelectById(id string) (*game.GameCore, error) {
	var data Game

	tx := repo.db.Where("id = ?", id).First(&data)

	if tx.RowsAffected == 0 {
		return &game.GameCore{}, errors.New("invalid id")
	}

	if tx.Error != nil {
		return &game.GameCore{}, tx.Error
	}

	return &game.GameCore{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Stock:       data.Stock,
		Discount:    data.Discount,
		Genre:       data.Genre,
		Publisher:   data.Publisher,
		ReleaseDate: data.ReleaseDate,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}, nil
}

// Insert implements game.DataInterface.
func (repo *gameRepository) Insert(data game.GameCore) (*game.GameCore, error) {
	var input = Game{
		ID:          uuid.New().String(),
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Stock:       data.Stock,
		Discount:    data.Discount,
		Genre:       data.Genre,
		Publisher:   data.Publisher,
		ReleaseDate: data.ReleaseDate,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
	tx := repo.db.Create(&input)

	if tx.Error != nil {
		return &game.GameCore{}, tx.Error
	}
	return &game.GameCore{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
		Discount:    input.Discount,
		Genre:       input.Genre,
		Publisher:   input.Publisher,
		ReleaseDate: input.ReleaseDate,
		CreatedAt:   input.CreatedAt,
		UpdatedAt:   input.UpdatedAt,
	}, nil
}

// Update implements game.DataInterface.
func (repo *gameRepository) Update(id string, data game.GameCore) error {
	tx := repo.db.Model(&Game{
		ID: id,
	}).Updates(Game{
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Stock:       data.Stock,
		Discount:    data.Discount,
		Genre:       data.Genre,
		Publisher:   data.Publisher,
		ReleaseDate: data.ReleaseDate,
	})

	if tx.RowsAffected == 0 {
		return errors.New("invalid id")
	}
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements game.DataInterface.
func (repo *gameRepository) Delete(id string) error {
	tx := repo.db.Where("id = ?", id).Delete(&Game{})

	if tx.RowsAffected == 0 {
		return errors.New("invalid id")
	}
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func New(db *gorm.DB) game.DataInterface {
	return &gameRepository{
		db: db,
	}
}
