package router

import (
	userHandler "mini_project/features/user/handler"
	userRepository "mini_project/features/user/repository"
	userUsecase "mini_project/features/user/usecase"

	adminHandler "mini_project/features/admin/handler"
	adminRepository "mini_project/features/admin/repository"
	adminUsecase "mini_project/features/admin/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userRepository := userRepository.New(db)
	userUsecase := userUsecase.New(userRepository)
	userController := userHandler.New(userUsecase)

	adminRepository := adminRepository.New(db)
	adminUsecase := adminUsecase.New(adminRepository)
	adminController := adminHandler.New(adminUsecase)

	user := e.Group("/user")
	user.POST("/register", userController.CreateUser)
	user.POST("/login", userController.UserLogin)

	admin := e.Group("/admin")
	admin.POST("/register", adminController.CreateUser)
	admin.POST("/login", adminController.UserLogin)
}
