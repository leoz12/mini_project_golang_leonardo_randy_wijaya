package transactionRepository

import (
	"errors"
	gameRepository "mini_project/features/game/repository"
	"mini_project/features/transaction"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

// SelectAll implements transaction.DataInterface.
func (repo *transactionRepository) SelectAll(userId string, role string) ([]transaction.Core, error) {
	var transactions []Transaction

	var tx *gorm.DB

	if role == "user" {
		tx = repo.db.Where("user_id = ?", userId).Preload("User").Find(&transactions)
	} else {
		tx = repo.db.Preload("User").Find(&transactions)

	}

	var transactionsCore []transaction.Core

	if tx.Error != nil {
		return transactionsCore, tx.Error
	}
	for _, val := range transactions {
		transactionsCore = append(transactionsCore, transaction.Core{
			Id:              val.ID,
			UserId:          val.UserId,
			User:            val.User,
			GameId:          val.GameId,
			GameName:        val.GameName,
			GameDescription: val.GameDescription,
			Quantity:        val.Quantity,
			Price:           val.Price,
			Discount:        val.Discount,
			CreatedAt:       val.CreatedAt,
			UpdatedAt:       val.UpdatedAt,
		})
	}

	return transactionsCore, nil
}

// SelectById implements transaction.DataInterface.
func (repo *transactionRepository) SelectById(id string) (transaction.Core, error) {
	var data Transaction

	tx := repo.db.Where("id = ?", id).First(&data)

	if tx.RowsAffected == 0 {
		return transaction.Core{}, errors.New("invalid id")
	}

	if tx.Error != nil {
		return transaction.Core{}, tx.Error
	}

	return transaction.Core{
		Id:              data.ID,
		UserId:          data.UserId,
		User:            data.User,
		GameId:          data.GameId,
		GameName:        data.GameName,
		GameDescription: data.GameDescription,
		Quantity:        data.Quantity,
		Price:           data.Price,
		Discount:        data.Discount,
		CreatedAt:       data.CreatedAt,
		UpdatedAt:       data.UpdatedAt,
	}, nil
}

// Insert implements transaction.DataInterface.
func (repo *transactionRepository) Insert(data transaction.Core) (transaction.Core, error) {
	var Game gameRepository.Game
	txGame := repo.db.Where("id = ?", data.GameId).First(&Game)

	if txGame.RowsAffected == 0 {
		return transaction.Core{}, errors.New("invalid game id")
	}

	if data.Quantity > Game.Stock {
		return transaction.Core{}, errors.New("insufficient stock, please enter the appropriate quantity")
	}
	var input = Transaction{
		ID:              uuid.New().String(),
		UserId:          data.UserId,
		GameId:          Game.ID,
		GameName:        Game.Name,
		GameDescription: Game.Description,
		Price:           Game.Price,
		Quantity:        data.Quantity,
		Discount:        Game.Discount,
	}
	tx := repo.db.Create(&input)
	if tx.Error != nil {
		return transaction.Core{}, tx.Error
	}

	txStock := repo.db.Model(&gameRepository.Game{}).Where("id = ?", Game.ID).Update("stock", Game.Stock-data.Quantity)

	if txStock.Error != nil {
		return transaction.Core{}, txStock.Error
	}

	return transaction.Core{
		Id:              input.ID,
		UserId:          input.UserId,
		GameId:          input.ID,
		GameName:        input.GameName,
		GameDescription: input.GameDescription,
		Price:           input.Price,
		Quantity:        input.Quantity,
		Discount:        input.Discount,
		CreatedAt:       input.CreatedAt,
		UpdatedAt:       input.UpdatedAt,
	}, nil
}

func New(db *gorm.DB) transaction.DataInterface {
	return &transactionRepository{
		db: db,
	}
}
