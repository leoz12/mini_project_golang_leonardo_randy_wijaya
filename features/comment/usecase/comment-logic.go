package commentUseCase

import (
	"errors"
	"mini_project/features/comment"
)

type commentUsecase struct {
	commentRepository comment.DataInterface
}

// GetAll implements comment.UseCaseInterface.
func (uc *commentUsecase) GetAll(gameId string) ([]comment.Core, error) {
	resp, err := uc.commentRepository.SelectAll(gameId)

	return resp, err
}

// GetById implements comment.UseCaseInterface.
func (uc *commentUsecase) GetById(id string) (comment.Core, error) {
	resp, err := uc.commentRepository.SelectById(id)

	if id == "" {
		return comment.Core{}, errors.New("id is required")
	}

	return resp, err
}

// Insert implements comment.UseCaseInterface.
func (uc *commentUsecase) Insert(role string, data comment.Core) (comment.Core, error) {
	if data.Comment == "" {
		return comment.Core{}, errors.New("comment is required")
	}
	response, err := uc.commentRepository.Insert(role, data)

	return response, err
}

// Update implements comment.UseCaseInterface.
func (uc *commentUsecase) Update(role string, data comment.Core) error {
	if data.Comment == "" {
		return errors.New("comment is required")
	}
	err := uc.commentRepository.Update(role, data)

	return err
}

// Delete implements comment.UseCaseInterface.
func (uc *commentUsecase) Delete(role string, data comment.Core) error {
	if data.Id == "" {
		return errors.New("id is required")
	}
	err := uc.commentRepository.Delete(role, data)
	return err
}

func New(commentRepo comment.DataInterface) comment.UseCaseInterface {
	return &commentUsecase{
		commentRepository: commentRepo,
	}
}
