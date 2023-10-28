package handler

import (
	"mini_project/app/middlewares"
	"mini_project/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadImageController(c echo.Context) error {

	tokenData := middlewares.ExtractToken(c)

	if tokenData.Id == "" {
		return c.JSON(http.StatusUnauthorized, helpers.FailedResponse("unauthorized"))
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error To upload Image")
	}

	resp, err := helpers.UploadImage(file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error To upload Image")
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "image upload success",
		"data":    resp,
	})
}
