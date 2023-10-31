package wishlistUseCase

import (
	"errors"
	"mini_project/features/wishlist"
)

type wishlistUsecase struct {
	wishlistRepository wishlist.DataInterface
}

func New(wishlistRepo wishlist.DataInterface) wishlist.UseCaseInterface {
	return &wishlistUsecase{
		wishlistRepository: wishlistRepo,
	}
}

func (uc *wishlistUsecase) GetAll(userId string) ([]wishlist.Core, error) {
	resp, err := uc.wishlistRepository.SelectAll(userId)

	return resp, err
}

func (uc *wishlistUsecase) Insert(data wishlist.Core) (wishlist.Core, error) {
	if data.GameId == "" {
		return wishlist.Core{}, errors.New("game id is required")
	}
	response, err := uc.wishlistRepository.Insert(data)

	return response, err
}

func (uc *wishlistUsecase) Delete(id string, userId string) error {
	if id == "" {
		return errors.New("id is required")
	}
	data, errGet := uc.wishlistRepository.SelectById(id)
	if errGet != nil {
		return errGet
	}
	if data.UserId != userId {
		return errors.New("unauthorized")
	}
	err := uc.wishlistRepository.Delete(id)
	return err
}
