package usecase

import (
	"errors"
	"mini_project/app/middlewares"
	"mini_project/features/admin"
	"mini_project/utils/helpers"
)

type adminUseCase struct {
	adminRepository admin.DataInterface
}

func (uc *adminUseCase) Create(data admin.AdminCore) error {
	if data.Email == "" || data.Password == "" {
		return errors.New("[validation] error. email dan password harus diisi")
	}

	hashPassword, errHash := helpers.HashPassword(data.Password)

	data.Password = hashPassword

	if errHash != nil {
		return errors.New(errHash.Error())
	}

	err := uc.adminRepository.Insert(data)
	return err
}

func (uc *adminUseCase) Login(data admin.LoginCore) (string, error) {

	if data.Email == "" || data.Password == "" {
		return "", errors.New("[validation] error. email dan password harus diisi")
	}

	dataUser, err := uc.adminRepository.CheckByEmail(data.Email)

	if err != nil {
		return "", err
	}

	if helpers.CheckPasswordHash(dataUser.Password, data.Password) {
		token, errToken := middlewares.CreateToken(dataUser.ID, "admin")

		if errToken != nil {
			return "", errToken
		}
		return token, nil
	} else {
		return "", errors.New("invalid email or password")
	}
}

func AdminUseCase(adminRepo admin.DataInterface) admin.UseCaseInterface {
	return &adminUseCase{
		adminRepository: adminRepo,
	}
}
