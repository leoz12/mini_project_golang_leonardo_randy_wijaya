package userUsecase

import (
	"errors"
	"mini_project/app/middlewares"
	"mini_project/features/user"
	"mini_project/utils/helpers"
)

type userUsecase struct {
	userRepository user.DataInterface
}

func (uc *userUsecase) Register(data user.UserCore) error {
	if data.Email == "" || data.Password == "" {
		return errors.New("[validation] error. email dan password harus diisi")
	}

	hashPassword, errHash := helpers.HashPassword(data.Password)

	data.Password = hashPassword

	if errHash != nil {
		return errors.New(errHash.Error())
	}

	err := uc.userRepository.Insert(data)
	return err
}

func (uc *userUsecase) Login(data user.LoginCore) (string, error) {

	if data.Email == "" || data.Password == "" {
		return "", errors.New("[validation] error. email dan password harus diisi")
	}

	dataUser, err := uc.userRepository.CheckByEmail(data.Email)

	if err != nil {
		return "", err
	}

	if helpers.CheckPasswordHash(dataUser.Password, data.Password) {
		token, errToken := middlewares.CreateToken(dataUser.ID, "user")

		if errToken != nil {
			return "", errToken
		}
		return token, nil
	} else {
		return "", errors.New("invalid email or password")
	}
}

func New(userRepo user.DataInterface) user.UseCaseInterface {
	return &userUsecase{
		userRepository: userRepo,
	}
}
