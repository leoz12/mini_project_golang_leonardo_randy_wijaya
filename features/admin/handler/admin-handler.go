package adminHandler

import (
	"errors"
	"mini_project/app/middlewares"
	"mini_project/features/admin"
	"mini_project/utils/helpers"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type adminController struct {
	adminUseCase admin.UseCaseInterface
}

func New(userUC admin.UseCaseInterface) *adminController {
	return &adminController{
		adminUseCase: userUC,
	}
}

func (handler *adminController) GetAllAdmin(c echo.Context) error {

	tokenData := middlewares.ExtractToken(c)

	if tokenData.Role != "admin" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	resp, err := handler.adminUseCase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	var responses = []AdminResponse{}

	for _, val := range resp {
		responses = append(responses, CoreToResponse(val))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success get all data", responses))
}

func (handler *adminController) CreateUser(c echo.Context) error {
	input := new(AdminRegisterRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error bind data"+errBind.Error()))
	}
	errValidations := helpers.ReqeustValidator(input)
	if len(errValidations) > 0 {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": errValidations,
		})
	}

	data := admin.Core{
		Email:    input.Email,
		Password: input.Password,
	}
	var mysqlErr *mysql.MySQLError

	err := handler.adminUseCase.Create(data)
	if err != nil {
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse("the email has already been taken"))
		}

		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusCreated, helpers.SuccessResponse("success insert data"))
}

func (handler *adminController) UserLogin(c echo.Context) error {
	input := new(AdminLoginRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("error bind data"+errBind.Error()))
	}

	errValidations := helpers.ReqeustValidator(input)
	if len(errValidations) > 0 {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": errValidations,
		})
	}

	data := admin.Core{
		Email:    input.Email,
		Password: input.Password,
	}
	token, err := handler.adminUseCase.Login(data)
	if err == gorm.ErrRecordNotFound {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse("invalid email or password"))
	} else if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]any{
		"token":   token,
		"message": "Login Success",
	})
}
