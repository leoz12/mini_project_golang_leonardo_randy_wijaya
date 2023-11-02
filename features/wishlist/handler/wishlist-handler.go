package wishlistHandler

import (
	"mini_project/app/configs"
	"mini_project/app/middlewares"
	"mini_project/features/wishlist"
	"mini_project/utils/helpers"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type wishlistController struct {
	wishlistUsecase wishlist.UseCaseInterface
}

func New(wishlistUC wishlist.UseCaseInterface) *wishlistController {
	return &wishlistController{
		wishlistUsecase: wishlistUC,
	}
}

func (handler *wishlistController) GetWishlists(c echo.Context) error {

	tokenData := middlewares.ExtractToken(c)

	if tokenData.Id == "" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	resp, err := handler.wishlistUsecase.GetAll(tokenData.Id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	var responses = []WishlistResponse{}

	for _, val := range resp {
		responses = append(responses, GetWishlistReponse(val))
	}
	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("success get all data", responses))
}

func (handler *wishlistController) CreateWishlist(c echo.Context) error {
	tokenData := middlewares.ExtractToken(c)

	if tokenData.Role != configs.UserRole.Admin {
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

	data := wishlist.Core{
		UserId: tokenData.Id,
		GameId: input.GameId,
	}
	_, err := handler.wishlistUsecase.Insert(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessResponse("success create wishlist"))
}

func (handler *wishlistController) DeleteWishlist(c echo.Context) error {
	tokenData := middlewares.ExtractToken(c)

	if tokenData.Role != configs.UserRole.Admin {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}
	id := c.Param("id")

	err := handler.wishlistUsecase.Delete(id, tokenData.Id)

	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			return c.JSON(http.StatusBadRequest, helpers.FailedResponse(err.Error()))
		} else if strings.Contains(err.Error(), "unauthorized") {
			return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
		}
		return c.JSON(http.StatusInternalServerError, helpers.FailedResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, helpers.SuccessResponse("success delete data"))
}
