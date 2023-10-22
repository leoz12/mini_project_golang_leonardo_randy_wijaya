package genreHandler

import (
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

func (handler *genreController) GetAllGenre(c echo.Context) error {

	resp, err := handler.genreUsecase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	var responses = []GenreDataResponse{}

	for _, val := range resp {
		responses = append(responses, GenreDataResponse{
			Id:        val.ID,
			Name:      val.Name,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success get all data", responses))
}

func (handler *genreController) CreateGenre(c echo.Context) error {
	input := new(GenreCreateRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error bind data",
		})
	}
	data := genre.GenreCore{
		Name: input.Name,
	}
	resp, err := handler.genreUsecase.Insert(data)

	if err != nil {
		if strings.Contains(err.Error(), "required") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success create genre", GenreDataResponse{
		Id:        resp.ID,
		Name:      resp.Name,
		CreatedAt: resp.CreatedAt,
		UpdatedAt: resp.UpdatedAt,
	}))
}
func (handler *genreController) UpdateGenre(c echo.Context) error {
	input := new(GenreCreateRequest)
	errBind := c.Bind(&input)
	id := c.Param("id")

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "error bind data",
		})
	}
	data := genre.GenreCore{
		Name: input.Name,
	}
	err := handler.genreUsecase.Update(id, data)

	if err != nil {
		if strings.Contains(err.Error(), "required") || strings.Contains(err.Error(), "invalid") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessResponse("success update genre"))
}

func (handler *genreController) DeleteGenre(c echo.Context) error {
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
