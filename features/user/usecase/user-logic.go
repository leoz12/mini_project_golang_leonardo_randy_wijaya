package userUsecase

import (
	"errors"
	"mini_project/app/middlewares"
	"mini_project/features/user"
	"mini_project/utils/helpers"
)

type userUseCase struct {
	userRepository user.DataInterface
}

func (uc *userUseCase) GetAll() ([]user.Core, error) {
	resp, err := uc.userRepository.SelectAll()

	return resp, err
}

func (uc *userUseCase) Register(data user.Core) error {
	if data.Email == "" || data.Password == "" {
		return errors.New("email and password are required")
	}

	hashPassword, errHash := helpers.HashPassword(data.Password)

	data.Password = hashPassword

	if errHash != nil {
		return errors.New(errHash.Error())
	}

	err := uc.userRepository.Insert(data)
	return err
}

func (uc *userUseCase) Login(data user.Core) (string, error) {

	if data.Email == "" || data.Password == "" {
		return "", errors.New("email and password are required")
	}

	dataUser, err := uc.userRepository.CheckByEmail(data.Email)

	if err != nil {
		return "", err
	}

	if helpers.CheckPasswordHash(dataUser.Password, data.Password) {
		token, errToken := middlewares.CreateToken(dataUser.Id, "user")

		if errToken != nil {
			return "", errToken
		}
		return token, nil
	} else {
		return "", errors.New("invalid email or password")
	}
}

func New(userRepo user.DataInterface) user.UseCaseInterface {
	return &userUseCase{
		userRepository: userRepo,
	}
}
