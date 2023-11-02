package userHandler

import (
	"errors"
	"mini_project/app/configs"
	"mini_project/app/middlewares"
	"mini_project/features/user"
	"mini_project/utils/helpers"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type userController struct {
	userUseCase user.UseCaseInterface
}

func New(userUC user.UseCaseInterface) *userController {
	return &userController{
		userUseCase: userUC,
	}
}

func (handler *userController) GetAllUser(c echo.Context) error {

	tokenData := middlewares.ExtractToken(c)

	if tokenData.Role != configs.UserRole.Admin {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	resp, err := handler.userUseCase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	var responses = []UserResponse{}

	for _, val := range resp {
		responses = append(responses, CoreToResponse(val))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success get all data", responses))
}

func (handler *userController) CreateUser(c echo.Context) error {
	input := new(UserRegisterRequest)
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
	data := user.Core{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	var mysqlErr *mysql.MySQLError

	err := handler.userUseCase.Register(data)
	if err != nil {
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse("the email has already been taken"))
		}

		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusCreated, helpers.SuccessResponse("success insert data"))
}

func (handler *userController) UserLogin(c echo.Context) error {
	input := new(UserLoginRequest)
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

	data := user.Core{
		Email:    input.Email,
		Password: input.Password,
	}
	token, err := handler.userUseCase.Login(data)
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
