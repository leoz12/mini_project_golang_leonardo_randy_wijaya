package gameHandler

import (
	"mini_project/app/middlewares"
	"mini_project/features/game"
	"mini_project/utils/helpers"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type gameController struct {
	gameUsecase game.UseCaseInterface
}

func New(gameUC game.UseCaseInterface) *gameController {
	return &gameController{
		gameUsecase: gameUC,
	}
}

func (handler *gameController) GetAllGame(c echo.Context) error {

	tokenData := middlewares.ExtractToken(c)

	if tokenData.Id == "" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	resp, err := handler.gameUsecase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	var responses = []GameLiteResponse{}

	for _, val := range resp {
		responses = append(responses, GetGameLiteResponse(val))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success get all data", responses))
}

func (handler *gameController) GetById(c echo.Context) error {
	tokenData := middlewares.ExtractToken(c)

	if tokenData.Id == "" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	id := c.Param("id")
	resp, err := handler.gameUsecase.GetById(id)

	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success get data", GetGameResponse(resp)))
}

func (handler *gameController) CreateGame(c echo.Context) error {
	tokenData := middlewares.ExtractToken(c)

	if tokenData.Role != "admin" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	input := new(CreateRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errBind,
		})
	}
	// Define the format layout that matches the input date string
	layout := "02-01-2006"

	parsedTime, errTime := time.Parse(layout, input.ReleaseDate)
	if errTime != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(errTime.Error()))
	}

	data := game.GameCore{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
		Discount:    input.Discount,
		Genre:       input.Genre,
		Publisher:   input.Publisher,
		ReleaseDate: parsedTime,
	}
	resp, err := handler.gameUsecase.Insert(data)

	if err != nil {
		if strings.Contains(err.Error(), "required") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success create genre", GetGameResponse(resp)))
}

func (handler *gameController) UpdateGame(c echo.Context) error {
	tokenData := middlewares.ExtractToken(c)

	if tokenData.Role != "admin" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	input := new(CreateRequest)
	errBind := c.Bind(&input)
	id := c.Param("id")

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": errBind,
		})
	}

	// Define the format layout that matches the input date string
	layout := "02-01-2006"

	parsedTime, errTime := time.Parse(layout, input.ReleaseDate)
	if errTime != nil {
		return c.JSON(http.StatusBadRequest, helpers.FailedResponse(errTime.Error()))
	}

	data := game.GameCore{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
		Discount:    input.Discount,
		Genre:       input.Genre,
		Publisher:   input.Publisher,
		ReleaseDate: parsedTime,
	}
	err := handler.gameUsecase.Update(id, data)

	if err != nil {
		if strings.Contains(err.Error(), "required") || strings.Contains(err.Error(), "invalid") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessResponse("success update data"))
}

func (handler *gameController) DeleteGame(c echo.Context) error {
	tokenData := middlewares.ExtractToken(c)

	if tokenData.Role != "admin" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}
	id := c.Param("id")

	err := handler.gameUsecase.Delete(id)

	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessResponse("success delete data"))
}
