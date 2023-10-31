package wishlistRepository

import (
	"errors"
	gameRepository "mini_project/features/game/repository"
	"mini_project/features/wishlist"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type wishlistRepository struct {
	db *gorm.DB
}

// SelectAll implements wishlist.DataInterface.
func (repo *wishlistRepository) SelectAll(userId string) ([]wishlist.Core, error) {
	var wishlists []Wishlist
	tx := repo.db.Where("user_id = ?", userId).Preload("Game").Preload("Game.Genres").Find(&wishlists)

	if tx.Error != nil {
		return nil, tx.Error
	}

	var wishlistsCore []wishlist.Core
	for _, val := range wishlists {
		wishlistsCore = append(wishlistsCore, ModelToCore(val))
	}
	return wishlistsCore, nil
}

// SelectById implements wishlist.DataInterface.
func (repo *wishlistRepository) SelectById(id string) (wishlist.Core, error) {
	var data Wishlist
	tx := repo.db.Where("id = ?", id).First(&data)

	if tx.RowsAffected == 0 {
		return wishlist.Core{}, errors.New("invalid id")
	}

	if tx.Error != nil {
		return wishlist.Core{}, tx.Error
	}
	return ModelToCore(data), nil
}

// Insert implements wishlist.DataInterface.
func (repo *wishlistRepository) Insert(data wishlist.Core) (wishlist.Core, error) {
	var Game gameRepository.Game
	txGame := repo.db.Where("id = ?", data.GameId).First(&Game)

	if txGame.RowsAffected == 0 {
		return wishlist.Core{}, errors.New("invalid game id")
	}

	var input = Wishlist{
		ID:     uuid.New().String(),
		UserID: data.UserId,
		GameID: data.GameId,
	}
	tx := repo.db.Create(&input)
	input.Game = Game
	if tx.Error != nil {
		return wishlist.Core{}, tx.Error
	}
	return ModelToCore(input), nil
}

// Delete implements wishlist.DataInterface.
func (repo *wishlistRepository) Delete(id string) error {
	tx := repo.db.Where("id = ?", id).Delete(&Wishlist{})

	if tx.RowsAffected == 0 {
		return errors.New("invalid id")
	}
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func New(db *gorm.DB) wishlist.DataInterface {
	return &wishlistRepository{
		db: db,
	}
}
