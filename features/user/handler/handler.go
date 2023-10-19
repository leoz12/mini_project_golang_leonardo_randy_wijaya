package handler

import (
	"mini_project/features/user"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase user.UseCaseInterface
}

func UserHandler(userUC user.UseCaseInterface) *UserController {
	return &UserController{
		userUsecase: userUC,
	}
}

func (handler *UserController) CreateUser(c echo.Context) error {
	input := new(UserRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := user.UserCore{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	err := handler.userUsecase.Create(data)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": err.Error(),
			})

		}
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": "error insert data. " + err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]any{
		"message": "success insert data",
	})
}

func (handler *UserController) UserLogin(c echo.Context) error {
	input := new(UserLoginRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := user.LoginCore{
		Email:    input.Email,
		Password: input.Password,
	}
	token, err := handler.userUsecase.Login(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})

	}

	return c.JSON(http.StatusOK, map[string]any{
		"token":   token,
		"message": "Login Success",
	})
}
