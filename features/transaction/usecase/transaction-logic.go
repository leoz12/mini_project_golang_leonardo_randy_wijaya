package transactionUseCase

import (
	"errors"
	"mini_project/features/transaction"
)

type transactionUsecase struct {
	transactionRepository transaction.DataInterface
}

func New(transactionRepo transaction.DataInterface) transaction.UseCaseInterface {
	return &transactionUsecase{
		transactionRepository: transactionRepo,
	}
}

// GetAll implements transaction.UseCaseInterface.
func (uc *transactionUsecase) GetAll(userId string, role string) ([]transaction.Core, error) {
	resp, err := uc.transactionRepository.SelectAll(userId, role)

	return resp, err
}

// GetById implements transaction.UseCaseInterface.
func (uc *transactionUsecase) GetById(id string) (transaction.Core, error) {
	resp, err := uc.transactionRepository.SelectById(id)

	if id == "" {
		return transaction.Core{}, errors.New("id is required")
	}

	return resp, err
}

// Insert implements transaction.UseCaseInterface.
func (uc *transactionUsecase) Insert(data transaction.Core) (transaction.Core, error) {
	// if data.Name == "" {
	// 	return nil, errors.New("name is required")
	// }
	response, err := uc.transactionRepository.Insert(data)

	return response, err
}
