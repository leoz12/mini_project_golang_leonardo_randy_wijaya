package adminUsecase

import (
	"errors"
	"mini_project/app/configs"
	"mini_project/app/middlewares"
	"mini_project/features/admin"
	"mini_project/utils/helpers"
)

type adminUseCase struct {
	adminRepository admin.DataInterface
}

func (uc *adminUseCase) GetAll() ([]admin.Core, error) {
	resp, err := uc.adminRepository.SelectAll()

	return resp, err
}

func (uc *adminUseCase) Create(data admin.Core) error {

	hashPassword, errHash := helpers.HashPassword(data.Password)

	data.Password = hashPassword

	if errHash != nil {
		return errors.New(errHash.Error())
	}

	err := uc.adminRepository.Insert(data)
	return err
}

func (uc *adminUseCase) Login(data admin.Core) (string, error) {

	dataUser, err := uc.adminRepository.CheckByEmail(data.Email)

	if err != nil {
		return "", err
	}

	if helpers.CheckPasswordHash(dataUser.Password, data.Password) {
		token, errToken := middlewares.CreateToken(dataUser.Id, configs.UserRole.Admin)

		if errToken != nil {
			return "", errToken
		}
		return token, nil
	} else {
		return "", errors.New("invalid email or password")
	}
}

func New(adminRepo admin.DataInterface) admin.UseCaseInterface {
	return &adminUseCase{
		adminRepository: adminRepo,
	}
}
