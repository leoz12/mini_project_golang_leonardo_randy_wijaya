package genreUseCase

import (
	"errors"
	"mini_project/features/genre"
)

type genreUsecase struct {
	genreRepository genre.DataInterface
}

// GetAll implements genre.UseCaseInterface.
func (uc *genreUsecase) GetAll() ([]genre.Core, error) {
	resp, err := uc.genreRepository.SelectAll()

	return resp, err
}

// GetById implements genre.UseCaseInterface.
func (uc *genreUsecase) GetById(id string) (genre.Core, error) {
	resp, err := uc.genreRepository.SelectById(id)

	if id == "" {
		return genre.Core{}, errors.New("id is required")
	}

	return resp, err
}

// Insert implements genre.UseCaseInterface.
func (uc *genreUsecase) Insert(data genre.Core) (genre.Core, error) {

	if data.Name == "" {
		return genre.Core{}, errors.New("name is required")
	}
	response, err := uc.genreRepository.Insert(data)

	return response, err
}

// Update implements genre.UseCaseInterface.
func (uc *genreUsecase) Update(id string, data genre.Core) error {

	if data.Name == "" {
		return errors.New("name is required")
	}
	err := uc.genreRepository.Update(id, data)

	return err
}

// Delete implements genre.UseCaseInterface.
func (uc *genreUsecase) Delete(id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	err := uc.genreRepository.Delete(id)
	return err
}

func New(genreRepo genre.DataInterface) genre.UseCaseInterface {
	return &genreUsecase{
		genreRepository: genreRepo,
	}
}
