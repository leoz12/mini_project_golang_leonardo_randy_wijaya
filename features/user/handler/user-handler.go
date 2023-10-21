package handler

import (
	"errors"
	"mini_project/features/user"
	"net/http"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
	input := new(UserRegisterRequest)
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
	var mysqlErr *mysql.MySQLError

	err := handler.userUsecase.Create(data)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})

		} else if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "the email has already been taken",
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]any{
			"message": err.Error(),
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
	if err == gorm.ErrRecordNotFound {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid email or password",
		})
	} else if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})

	}

	return c.JSON(http.StatusOK, map[string]any{
		"token":   token,
		"message": "Login Success",
	})
}
