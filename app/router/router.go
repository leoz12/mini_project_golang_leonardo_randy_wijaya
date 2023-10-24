package router

import (
	"mini_project/app/middlewares"
	genreHandler "mini_project/features/genre/handler"
	genreRepository "mini_project/features/genre/repository"
	genreUseCase "mini_project/features/genre/usecase"
	userHandler "mini_project/features/user/handler"
	userRepository "mini_project/features/user/repository"
	userUsecase "mini_project/features/user/usecase"

	adminHandler "mini_project/features/admin/handler"
	adminRepository "mini_project/features/admin/repository"
	adminUsecase "mini_project/features/admin/usecase"

	gameHandler "mini_project/features/game/handler"
	gameRepository "mini_project/features/game/repository"
	gameUseCase "mini_project/features/game/usecase"

	wishlistHandler "mini_project/features/wishlist/handler"
	wishlistRepository "mini_project/features/wishlist/repository"
	wishlistUseCase "mini_project/features/wishlist/usecase"

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

	genreRepository := genreRepository.New(db)
	genreUsecase := genreUseCase.New(genreRepository)
	genreController := genreHandler.New(genreUsecase)

	gameRepository := gameRepository.New(db)
	gameUsecase := gameUseCase.New(gameRepository)
	gameController := gameHandler.New(gameUsecase)

	wishlistRepository := wishlistRepository.New(db)
	wishlistUsecase := wishlistUseCase.New(wishlistRepository)
	wishlistController := wishlistHandler.New(wishlistUsecase)

	user := e.Group("/user")
	user.POST("/register", userController.CreateUser)
	user.POST("/login", userController.UserLogin)

	admin := e.Group("/admin")
	admin.POST("/register", adminController.CreateUser)
	admin.POST("/login", adminController.UserLogin)

	genre := e.Group("/genre", middlewares.JWTMiddleware())
	genre.GET("", genreController.GetAllGenre)
	genre.POST("", genreController.CreateGenre)
	genre.PUT("/:id", genreController.UpdateGenre)
	genre.DELETE("/:id", genreController.DeleteGenre)

	game := e.Group("/game", middlewares.JWTMiddleware())
	game.GET("", gameController.GetAllGame)
	game.GET("/:id", gameController.GetById)
	game.POST("", gameController.CreateGame)
	game.PUT("/:id", gameController.UpdateGame)
	game.DELETE("/:id", gameController.DeleteGame)

	wishlists := e.Group("/wishlists", middlewares.JWTMiddleware())
	wishlists.GET("", wishlistController.GetWishlists)
	wishlists.POST("", wishlistController.CreateWishlist)
	wishlists.DELETE("/:id", wishlistController.DeleteWishlist)
}
