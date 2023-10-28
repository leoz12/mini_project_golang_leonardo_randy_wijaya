package router

import (
	"mini_project/app/middlewares"
	genreHandler "mini_project/features/genre/handler"
	genreRepository "mini_project/features/genre/repository"
	genreUseCase "mini_project/features/genre/usecase"
	"mini_project/features/upload-image/handler"
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

	transactionHandler "mini_project/features/transaction/handler"
	transactionRepository "mini_project/features/transaction/repository"
	transactionUseCase "mini_project/features/transaction/usecase"

	commentHandler "mini_project/features/comment/handler"
	commentRepository "mini_project/features/comment/repository"
	commentUseCase "mini_project/features/comment/usecase"

	recommendationHandler "mini_project/features/recommendation/handler"
	recommendationRepository "mini_project/features/recommendation/repository"
	recommendationUseCase "mini_project/features/recommendation/usecase"

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

	transactionRepository := transactionRepository.New(db)
	transactionUsecase := transactionUseCase.New(transactionRepository)
	transactionController := transactionHandler.New(transactionUsecase)

	commentRepository := commentRepository.New(db)
	commentUsecase := commentUseCase.New(commentRepository)
	commentController := commentHandler.New(commentUsecase)

	recommendationRepository := recommendationRepository.New(db)
	recommendationUsecase := recommendationUseCase.New(recommendationRepository)
	recommendationController := recommendationHandler.New(recommendationUsecase)

	user := e.Group("/user")
	user.GET("", userController.GetAllUser, middlewares.JWTMiddleware())
	user.POST("/register", userController.CreateUser)
	user.POST("/login", userController.UserLogin)

	admin := e.Group("/admin")
	admin.GET("", adminController.GetAllAdmin, middlewares.JWTMiddleware())
	admin.POST("/register", adminController.CreateUser)
	admin.POST("/login", adminController.UserLogin)

	genre := e.Group("/genres", middlewares.JWTMiddleware())
	genre.GET("", genreController.GetAllGenres)
	genre.GET("/:id", genreController.GetGenreById)
	genre.POST("", genreController.CreateGenre)
	genre.PUT("/:id", genreController.UpdateGenre)
	genre.DELETE("/:id", genreController.DeleteGenre)

	game := e.Group("/games", middlewares.JWTMiddleware())
	game.GET("", gameController.GetAllGames)
	game.GET("/:id", gameController.GetGameById)
	game.POST("", gameController.CreateGame)
	game.PUT("/:id", gameController.UpdateGame)
	game.DELETE("/:id", gameController.DeleteGame)

	wishlists := e.Group("/wishlists", middlewares.JWTMiddleware())
	wishlists.GET("", wishlistController.GetWishlists)
	wishlists.POST("", wishlistController.CreateWishlist)
	wishlists.DELETE("/:id", wishlistController.DeleteWishlist)

	transactions := e.Group("/transactions", middlewares.JWTMiddleware())
	transactions.GET("", transactionController.GetAllTransactions)
	transactions.GET("/:id", transactionController.GetTransactionById)
	transactions.POST("", transactionController.CreateTransaction)

	comments := e.Group("/comments", middlewares.JWTMiddleware())
	comments.GET("/game/:gameId", commentController.GetAllComment)
	comments.GET("/:id", commentController.GetCommentById)
	comments.POST("", commentController.CreateComment)
	comments.PUT("/:id", commentController.UpdateComment)
	comments.DELETE("/:id", commentController.DeleteComment)

	recommendations := e.Group("/recommendations", middlewares.JWTMiddleware())
	recommendations.GET("/game/:id", recommendationController.Getrecommendation)
	recommendations.GET("", recommendationController.GetAllrecommendations)
	recommendations.GET("/:id", recommendationController.GetrecommendationById)
	recommendations.POST("", recommendationController.Createrecommendation)
	recommendations.PUT("/:id", recommendationController.Updaterecommendation)
	recommendations.DELETE("/:id", recommendationController.Deleterecommendation)

	e.POST("upload-image", handler.UploadImageController)
}
