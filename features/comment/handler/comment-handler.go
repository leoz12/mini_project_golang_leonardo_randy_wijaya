package commentHandler

import (
	"mini_project/app/middlewares"
	"mini_project/features/comment"
	"mini_project/utils/helpers"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type commentController struct {
	commentUsecase comment.UseCaseInterface
}

func New(commentUC comment.UseCaseInterface) *commentController {
	return &commentController{
		commentUsecase: commentUC,
	}
}

func (handler *commentController) GetAllComment(c echo.Context) error {

	tokenData := middlewares.ExtractToken(c)

	if tokenData.Id == "" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}
	gameId := c.Param("gameId")

	resp, err := handler.commentUsecase.GetAll(gameId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	var responses = []CommentResponse{}

	for _, val := range resp {
		responses = append(responses, CoreToResponse(val, tokenData.Id))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success get all data", responses))
}

func (handler *commentController) GetCommentById(c echo.Context) error {
	tokenData := middlewares.ExtractToken(c)

	if tokenData.Id == "" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	id := c.Param("id")
	resp, err := handler.commentUsecase.GetById(id)

	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success get data", CoreToResponse(resp, tokenData.Id)))
}

func (handler *commentController) CreateComment(c echo.Context) error {
	tokenData := middlewares.ExtractToken(c)

	if tokenData.Id == "" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	input := new(CreateRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errBind,
		})
	}

	data := comment.Core{
		UserId:  tokenData.Id,
		Comment: input.Comment,
		GameId:  input.GameId,
	}
	resp, err := handler.commentUsecase.Insert(tokenData.Role, data)

	if err != nil {
		if strings.Contains(err.Error(), "required") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success create comment", CoreToResponseLite(resp)))
}

func (handler *commentController) UpdateComment(c echo.Context) error {
	tokenData := middlewares.ExtractToken(c)

	if tokenData.Id == "" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	input := new(UpdateRequest)
	errBind := c.Bind(&input)
	id := c.Param("id")

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": errBind,
		})
	}
	data := comment.Core{
		Id:      id,
		UserId:  tokenData.Id,
		Comment: input.Comment,
	}
	err := handler.commentUsecase.Update(tokenData.Role, data)

	if err != nil {
		if strings.Contains(err.Error(), "required") || strings.Contains(err.Error(), "invalid") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessResponse("success update comment"))
}

func (handler *commentController) DeleteComment(c echo.Context) error {
	tokenData := middlewares.ExtractToken(c)

	if tokenData.Id == "" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}
	id := c.Param("id")

	err := handler.commentUsecase.Delete(tokenData.Role, comment.Core{
		Id:     id,
		UserId: tokenData.Id,
	})

	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessResponse("success delete comment"))
}
