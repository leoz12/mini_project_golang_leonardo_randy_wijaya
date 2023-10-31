package gameUseCase

import (
	"errors"
	"mini_project/features/game"
	"time"
)

type gameUsecase struct {
	gameRepository game.DataInterface
}

// GetAll implements game.UseCaseinterface.
func (uc *gameUsecase) GetAll(params game.GameParams) ([]game.Core, error) {
	resp, err := uc.gameRepository.SelectAll(params)

	return resp, err
}

// GetById implements game.UseCaseinterface.
func (uc *gameUsecase) GetById(id string, userId string) (game.Core, error) {
	resp, err := uc.gameRepository.SelectById(id, userId)

	if id == "" {
		return game.Core{}, errors.New("id is required")
	}

	return resp, err
}

// Insert implements game.UseCaseinterface.
func (uc *gameUsecase) Insert(data game.Core) (game.Core, error) {
	// Define the format layout that matches the input date string
	if len(data.Genres) < 1 {
		return game.Core{}, errors.New("genres must be filled in")
	}

	layout := "02-01-2006"

	parsedTime, errTime := time.Parse(layout, data.ReleaseDateString)
	if errTime != nil {
		return game.Core{}, errors.New("invalid time format, time format should be dd-mm-yyyy")
	}
	data.ReleaseDate = parsedTime
	response, err := uc.gameRepository.Insert(data)

	return response, err
}

// Update implements game.UseCaseinterface.
func (uc *gameUsecase) Update(id string, data game.Core) error {

	// Define the format layout that matches the input date string
	layout := "02-01-2006"

	if len(data.Genres) < 1 {
		return errors.New("genres must be filled in")
	}

	parsedTime, errTime := time.Parse(layout, data.ReleaseDateString)
	if errTime != nil {
		return errors.New("invalid time format, time format should be dd-mm-yyyy")
	}
	data.ReleaseDate = parsedTime
	err := uc.gameRepository.Update(id, data)

	return err
}

// Delete implements game.UseCaseinterface.
func (uc *gameUsecase) Delete(id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	err := uc.gameRepository.Delete(id)
	return err
}

func New(gameRepo game.DataInterface) game.UseCaseInterface {
	return &gameUsecase{
		gameRepository: gameRepo,
	}
}
