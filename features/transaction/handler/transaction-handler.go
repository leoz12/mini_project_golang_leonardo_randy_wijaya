package transactionHandler

import (
	"mini_project/app/middlewares"
	"mini_project/features/transaction"
	"mini_project/utils/helpers"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type transactionController struct {
	transactionUsecase transaction.UseCaseInterface
}

func New(transactionUC transaction.UseCaseInterface) *transactionController {
	return &transactionController{
		transactionUsecase: transactionUC,
	}
}

func (handler *transactionController) GetAllTransactions(c echo.Context) error {

	tokenData := middlewares.ExtractToken(c)

	if tokenData.Id == "" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	resp, err := handler.transactionUsecase.GetAll(tokenData.Id, tokenData.Role)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	var responses = []TransactionResponse{}

	for _, val := range resp {
		responses = append(responses, CoreToResponse(val))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success get all data", responses))
}

func (handler *transactionController) GetTransactionById(c echo.Context) error {
	tokenData := middlewares.ExtractToken(c)

	if tokenData.Id == "" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	id := c.Param("id")
	resp, err := handler.transactionUsecase.GetById(id)

	if resp.UserId != tokenData.Id && tokenData.Role == "user" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success get data", CoreToResponse(resp)))
}

func (handler *transactionController) CreateTransaction(c echo.Context) error {
	tokenData := middlewares.ExtractToken(c)

	if tokenData.Role != "user" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	input := new(CreateRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": errBind,
		})
	}
	errValidations := helpers.ReqeustValidator(input)
	if len(errValidations) > 0 {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": errValidations,
		})
	}
	data := transaction.Core{
		UserId:   tokenData.Id,
		GameId:   input.GameId,
		Quantity: input.Quantity,
	}
	resp, err := handler.transactionUsecase.Insert(data)

	if err != nil {
		if strings.Contains(err.Error(), "required") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success create genre", CoreToResponse(resp)))
}
