package recommendationHandler

import (
	"mini_project/app/middlewares"
	"mini_project/features/recommendation"
	"mini_project/utils/helpers"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type recommendationController struct {
	recommendationUsecase recommendation.UseCaseInterface
}

func New(recommendationUC recommendation.UseCaseInterface) *recommendationController {
	return &recommendationController{
		recommendationUsecase: recommendationUC,
	}
}

func (handler *recommendationController) Getrecommendation(c echo.Context) error {
	tokenData := middlewares.ExtractToken(c)

	if tokenData.Id == "" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	id := c.Param("id")
	resp, err := handler.recommendationUsecase.RecommendGame(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success get data", resp))
}

func (handler *recommendationController) GetAllrecommendations(c echo.Context) error {

	tokenData := middlewares.ExtractToken(c)

	if tokenData.Id == "" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	resp, err := handler.recommendationUsecase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	var responses = []RecommendationResponse{}

	for _, val := range resp {
		responses = append(responses, CoreToResponse(val))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success get all data", responses))
}

func (handler *recommendationController) GetrecommendationById(c echo.Context) error {
	tokenData := middlewares.ExtractToken(c)

	if tokenData.Id == "" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	id := c.Param("id")
	resp, err := handler.recommendationUsecase.GetById(id)

	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success get data", CoreToResponse(resp)))
}

func (handler *recommendationController) Createrecommendation(c echo.Context) error {
	tokenData := middlewares.ExtractToken(c)

	if tokenData.Role != "admin" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	input := new(CreateRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error bind data",
		})
	}
	data := recommendation.Core{
		Name: input.Name,
	}
	resp, err := handler.recommendationUsecase.Insert(data)

	if err != nil {
		if strings.Contains(err.Error(), "required") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success create data", CoreToResponse(resp)))
}

func (handler *recommendationController) Updaterecommendation(c echo.Context) error {
	tokenData := middlewares.ExtractToken(c)

	if tokenData.Role != "admin" {
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
	data := recommendation.Core{
		Name: input.Name,
	}
	err := handler.recommendationUsecase.Update(id, data)

	if err != nil {
		if strings.Contains(err.Error(), "required") || strings.Contains(err.Error(), "invalid") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessResponse("success update data"))
}

func (handler *recommendationController) Deleterecommendation(c echo.Context) error {
	tokenData := middlewares.ExtractToken(c)

	if tokenData.Role != "admin" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	id := c.Param("id")

	err := handler.recommendationUsecase.Delete(id)

	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessResponse("success delete data"))
}
