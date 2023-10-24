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

func (uc *wishlistUsecase) GetAll(userId string) ([]wishlist.WishlistCore, error) {
	resp, err := uc.wishlistRepository.SelectAll(userId)

	return resp, err
}

func (uc *wishlistUsecase) Insert(data wishlist.WishlistCore) (wishlist.WishlistCore, error) {
	if data.GameId == "" {
		return wishlist.WishlistCore{}, errors.New("game id is required")
	}
	response, err := uc.wishlistRepository.Insert(data)

	return response, err
}

func (uc *wishlistUsecase) Delete(id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	err := uc.wishlistRepository.Delete(id)
	return err
}
