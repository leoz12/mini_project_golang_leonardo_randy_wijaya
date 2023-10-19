package router

import (
	"mini_project/features/user/handler"
	"mini_project/features/user/repository"
	"mini_project/features/user/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userRepository := repository.UserDB(db)
	userUsecase := usecase.UserUseCase(userRepository)
	userController := handler.UserHandler(userUsecase)

	user := e.Group("/user")
	user.POST("/register", userController.CreateUser)
	user.POST("/login", userController.UserLogin)
}
