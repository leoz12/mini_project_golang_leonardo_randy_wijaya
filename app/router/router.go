package router

import (
	userHandler "mini_project/features/user/handler"
	userRepo "mini_project/features/user/repository"
	userCase "mini_project/features/user/usecase"

	adminHandler "mini_project/features/admin/handler"
	adminRepo "mini_project/features/admin/repository"
	adminCase "mini_project/features/admin/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userRepository := userRepo.UserDB(db)
	userUsecase := userCase.UserUseCase(userRepository)
	userController := userHandler.UserHandler(userUsecase)

	adminRepository := adminRepo.AdminDB(db)
	adminUsecase := adminCase.AdminUseCase(adminRepository)
	adminController := adminHandler.AdminHandler(adminUsecase)

	user := e.Group("/user")
	user.POST("/register", userController.CreateUser)
	user.POST("/login", userController.UserLogin)

	admin := e.Group("/admin")
	admin.POST("/register", adminController.CreateUser)
	admin.POST("/login", adminController.UserLogin)
}
