package genreHandler

import (
	"mini_project/app/configs"
	"mini_project/app/middlewares"
	"mini_project/features/genre"
	"mini_project/utils/helpers"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type genreController struct {
	genreUsecase genre.UseCaseInterface
}

func New(genreUC genre.UseCaseInterface) *genreController {
	return &genreController{
		genreUsecase: genreUC,
	}
}

func (handler *genreController) GetAllGenres(c echo.Context) error {

	tokenData := middlewares.ExtractToken(c)

	if tokenData.Id == "" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	resp, err := handler.genreUsecase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	var responses = []GenreResponse{}

	for _, val := range resp {
		responses = append(responses, CoreToResponse(val))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success get all data", responses))
}

func (handler *genreController) GetGenreById(c echo.Context) error {
	tokenData := middlewares.ExtractToken(c)

	if tokenData.Id == "" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	id := c.Param("id")
	resp, err := handler.genreUsecase.GetById(id)

	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success get data", CoreToResponse(resp)))
}

func (handler *genreController) CreateGenre(c echo.Context) error {
	tokenData := middlewares.ExtractToken(c)

	if tokenData.Role != configs.UserRole.Admin {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	input := new(CreateRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error bind data",
		})
	}
	errValidations := helpers.ReqeustValidator(input)
	if len(errValidations) > 0 {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": errValidations,
		})
	}
	data := genre.Core{
		Name: input.Name,
	}
	resp, err := handler.genreUsecase.Insert(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success create data", GenreResponse{
		Id:        resp.Id,
		Name:      resp.Name,
		CreatedAt: resp.CreatedAt,
		UpdatedAt: resp.UpdatedAt,
	}))
}

func (handler *genreController) UpdateGenre(c echo.Context) error {
	tokenData := middlewares.ExtractToken(c)

	if tokenData.Role != configs.UserRole.Admin {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	input := new(CreateRequest)
	errBind := c.Bind(&input)
	id := c.Param("id")

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}
	errValidations := helpers.ReqeustValidator(input)
	if len(errValidations) > 0 {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": errValidations,
		})
	}
	data := genre.Core{
		Name: input.Name,
	}
	err := handler.genreUsecase.Update(id, data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessResponse("success update data"))
}

func (handler *genreController) DeleteGenre(c echo.Context) error {
	tokenData := middlewares.ExtractToken(c)

	if tokenData.Role != configs.UserRole.Admin {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	id := c.Param("id")

	err := handler.genreUsecase.Delete(id)

	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessResponse("success delete data"))
}
