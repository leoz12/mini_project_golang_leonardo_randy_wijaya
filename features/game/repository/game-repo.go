package gameRepository

import (
	"errors"
	"mini_project/features/game"
	"mini_project/features/genre"
	genreRepository "mini_project/features/genre/repository"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type gameRepository struct {
	db *gorm.DB
}

// SelectAll implements game.DataInterface.
func (repo *gameRepository) SelectAll(params game.GameParams) ([]game.Core, error) {
	var games []Game
	var tx *gorm.DB
	if len(strings.Split(params.Genres, ",")) > 0 && params.Genres != "" {
		tx = repo.db.Preload("Genres").Joins("LEFT JOIN game_genres gg ON gg.game_id = games.id").
			Joins("LEFT JOIN genres g ON g.id = gg.genre_id").
			Where("games.name LIKE ? AND g.name IN (?)", "%"+params.Search+"%", strings.Split(params.Genres, ",")).Group("id").Find(&games)
	} else {
		tx = repo.db.Preload("Genres").Joins("LEFT JOIN game_genres gg ON gg.game_id = games.id").
			Joins("LEFT JOIN genres g ON g.id = gg.genre_id").
			Where("games.name LIKE ?", "%"+params.Search+"%").Group("id").Find(&games)
	}

	var gamesCore []game.Core

	if tx.Error != nil {
		return gamesCore, tx.Error
	}
	for _, val := range games {
		var genresCore []genre.Core
		for _, cur := range val.Genres {
			genresCore = append(genresCore, genre.Core{
				Id:   cur.ID,
				Name: cur.Name,
			})
		}
		gamesCore = append(gamesCore, game.Core{
			Id:          val.ID,
			Name:        val.Name,
			Description: val.Description,
			Price:       val.Price,
			Stock:       val.Stock,
			Discount:    val.Discount,
			Genres:      genresCore,
			ImageUrl:    val.ImageUrl,
			Publisher:   val.Publisher,
			Platform:    val.Platform,
			ReleaseDate: val.ReleaseDate,
			CreatedAt:   val.CreatedAt,
			UpdatedAt:   val.UpdatedAt,
		})
	}
	return gamesCore, nil
}

// SelectById implements game.DataInterface.
func (repo *gameRepository) SelectById(id string, userId string) (game.Core, error) {
	var data Game

	tx := repo.db.Preload("Genres").Where("id = ?", id).First(&data)

	if tx.RowsAffected == 0 {
		return game.Core{}, errors.New("invalid id")
	}

	if tx.Error != nil {
		return game.Core{}, tx.Error
	}

	var count int
	txTransactions := repo.db.Raw("SELECT COUNT(*) FROM transactions WHERE game_id = ? AND user_id = ?", id, userId).Scan(&count)

	if txTransactions.Error != nil {
		return game.Core{}, txTransactions.Error
	}
	var canComment = false

	if count > 0 {
		canComment = true
	}

	var genresCore []genre.Core
	for _, cur := range data.Genres {
		genresCore = append(genresCore, genre.Core{
			Id:   cur.ID,
			Name: cur.Name,
		})
	}

	return game.Core{
		Id:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Stock:       data.Stock,
		Discount:    data.Discount,
		Genres:      genresCore,
		ImageUrl:    data.ImageUrl,
		Publisher:   data.Publisher,
		Platform:    data.Platform,
		CanComment:  canComment,
		ReleaseDate: data.ReleaseDate,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}, nil
}

// Insert implements game.DataInterface.
func (repo *gameRepository) Insert(data game.Core) (game.Core, error) {
	var genres []genreRepository.Genre

	for _, val := range data.Genres {
		genres = append(genres, genreRepository.Genre{
			ID: val.Id,
		})
	}

	var input = Game{
		ID:          uuid.New().String(),
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Stock:       data.Stock,
		Discount:    data.Discount,
		Genres:      genres,
		ImageUrl:    data.ImageUrl,
		Publisher:   data.Publisher,
		ReleaseDate: data.ReleaseDate,
		Platform:    data.Platform,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
	tx := repo.db.Create(&input)

	if tx.Error != nil {
		return game.Core{}, tx.Error
	}
	return game.Core{
		Id:          input.ID,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
		Discount:    input.Discount,
		Publisher:   input.Publisher,
		ReleaseDate: input.ReleaseDate,
		ImageUrl:    input.ImageUrl,
		Platform:    input.Platform,
		CreatedAt:   input.CreatedAt,
		UpdatedAt:   input.UpdatedAt,
	}, nil
}

// Update implements game.DataInterface.
func (repo *gameRepository) Update(id string, data game.Core) error {
	var genres []genreRepository.Genre

	for _, val := range data.Genres {
		genres = append(genres, genreRepository.Genre{
			ID: val.Id,
		})
	}
	tx := repo.db.Model(&Game{
		ID: id,
	}).Updates(Game{
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Stock:       data.Stock,
		Discount:    data.Discount,
		Genres:      genres,
		Publisher:   data.Publisher,
		ReleaseDate: data.ReleaseDate,
		ImageUrl:    data.ImageUrl,
		Platform:    data.Platform,
	})
	txGenre := repo.db.Model(&Game{
		ID: id,
	}).Association("Genres").Replace(genres)

	if tx.RowsAffected == 0 {
		return errors.New("invalid id")
	}
	if tx.Error != nil {
		return tx.Error
	}
	if txGenre != nil {
		return txGenre
	}
	return nil
}

// Delete implements game.DataInterface.
func (repo *gameRepository) Delete(id string) error {
	tx := repo.db.Where("id = ?", id).Delete(&Game{})

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
