package gameHandler

import (
	"mini_project/app/middlewares"
	"mini_project/features/game"
	"mini_project/features/genre"
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

func (handler *gameController) GetAllGames(c echo.Context) error {
	genres := c.QueryParam("genres")
	search := c.QueryParam("search")

	tokenData := middlewares.ExtractToken(c)

	if tokenData.Id == "" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	resp, err := handler.gameUsecase.GetAll(game.GameParams{
		Genres: genres,
		Search: search,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	var responses = []GameLiteResponse{}

	for _, val := range resp {
		responses = append(responses, CoreToLiteReponse(val))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success get all data", responses))
}

func (handler *gameController) GetGameById(c echo.Context) error {
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
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success get data", CoreToReponse(resp)))
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

	var genresCore []genre.Core
	for _, val := range input.Genres {
		genresCore = append(genresCore, genre.Core{
			Id: val,
		})
	}

	data := game.Core{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
		Discount:    input.Discount,
		Genres:      genresCore,
		Publisher:   input.Publisher,
		ImageUrl:    input.ImageUrl,
		Platform:    input.Platform,
		ReleaseDate: parsedTime,
	}
	resp, err := handler.gameUsecase.Insert(data)

	if err != nil {
		if strings.Contains(err.Error(), "required") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success create genre", CoreToCreateReponse(resp)))
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
	var genresCore []genre.Core

	for _, val := range input.Genres {
		genresCore = append(genresCore, genre.Core{
			Id: val,
		})
	}
	data := game.Core{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
		Discount:    input.Discount,
		Genres:      genresCore,
		Publisher:   input.Publisher,
		ReleaseDate: parsedTime,
		ImageUrl:    input.ImageUrl,
		Platform:    input.Platform,
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
