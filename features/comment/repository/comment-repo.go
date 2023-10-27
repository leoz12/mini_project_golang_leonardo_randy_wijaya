package commentRepository

import (
	"errors"
	"mini_project/features/comment"
	gameRepository "mini_project/features/game/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

// SelectAll implements comment.DataInterface.
func (repo *commentRepository) SelectAll(gameId string) ([]comment.Core, error) {
	var commments []Comment
	tx := repo.db.Preload("User").Where("game_id = ?", gameId).Find(&commments)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var wishlistsCore []comment.Core
	for _, val := range commments {
		wishlistsCore = append(wishlistsCore, ModelToCore(val))
	}
	return wishlistsCore, nil
}

// SelectById implements comment.DataInterface.
func (repo *commentRepository) SelectById(id string) (comment.Core, error) {
	var data Comment

	tx := repo.db.Where("id = ?", id).Preload("User").First(&data)

	if tx.RowsAffected == 0 {
		return comment.Core{}, errors.New("invalid id")
	}

	if tx.Error != nil {
		return comment.Core{}, tx.Error
	}
	return ModelToCore(data), nil
}

// Insert implements comment.DataInterface.
func (repo *commentRepository) Insert(role string, data comment.Core) (comment.Core, error) {
	var Game gameRepository.Game
	txGame := repo.db.Where("id = ?", data.GameId).First(&Game)

	if txGame.RowsAffected == 0 {
		return comment.Core{}, errors.New("invalid game id")
	}

	var input = Comment{
		ID:      uuid.New().String(),
		Comment: data.Comment,
		UserID:  data.UserId,
		GameID:  data.GameId,
	}

	tx := repo.db.Create(&input)
	if tx.Error != nil {
		return comment.Core{}, tx.Error
	}
	return ModelToCore(input), nil
}

// Update implements comment.DataInterface.
func (repo *commentRepository) Update(role string, data comment.Core) error {
	tx := repo.db.Model(&Comment{
		ID: data.Id,
	}).Updates(Comment{
		Comment: data.Comment,
	})

	if tx.RowsAffected == 0 {
		return errors.New("invalid id")
	}
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements comment.DataInterface.
func (repo *commentRepository) Delete(role string, data comment.Core) error {
	tx := repo.db.Where("id = ?", data.Id).Delete(&Comment{})

	if tx.RowsAffected == 0 {
		return errors.New("invalid id")
	}
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func New(db *gorm.DB) comment.DataInterface {
	return &commentRepository{
		db: db,
	}
}
