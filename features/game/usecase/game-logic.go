package gameUseCase

import (
	"errors"
	"mini_project/features/game"
)

type gameUsecase struct {
	gameRepository game.DataInterface
}

// GetAll implements game.UseCaseinterface.
func (uc *gameUsecase) GetAll() ([]game.GameCore, error) {
	resp, err := uc.gameRepository.SelectAll()

	return resp, err
}

// GetById implements game.UseCaseinterface.
func (uc *gameUsecase) GetById(id string) (*game.GameCore, error) {
	resp, err := uc.gameRepository.SelectById(id)

	if id == "" {
		return &game.GameCore{}, errors.New("id is required")
	}

	return resp, err
}

// Insert implements game.UseCaseinterface.
func (uc *gameUsecase) Insert(data game.GameCore) (*game.GameCore, error) {
	if data.Name == "" {
		return nil, errors.New("name is required")
	}
	response, err := uc.gameRepository.Insert(data)

	return response, err
}

// Update implements game.UseCaseinterface.
func (uc *gameUsecase) Update(id string, data game.GameCore) error {

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
