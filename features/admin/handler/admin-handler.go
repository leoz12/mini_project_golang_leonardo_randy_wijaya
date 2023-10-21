package handler

import (
	"errors"
	"mini_project/features/admin"
	"net/http"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AdminController struct {
	adminUseCase admin.UseCaseInterface
}

func AdminHandler(userUC admin.UseCaseInterface) *AdminController {
	return &AdminController{
		adminUseCase: userUC,
	}
}

func (handler *AdminController) CreateUser(c echo.Context) error {
	input := new(AdminRegisterRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := admin.AdminCore{
		Email:    input.Email,
		Password: input.Password,
	}
	var mysqlErr *mysql.MySQLError

	err := handler.adminUseCase.Create(data)
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

func (handler *AdminController) UserLogin(c echo.Context) error {
	input := new(AdminLoginRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}

	data := admin.LoginCore{
		Email:    input.Email,
		Password: input.Password,
	}
	token, err := handler.adminUseCase.Login(data)
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
