package migration

import (
	adminRepository "mini_project/features/admin/repository"
	commentRepository "mini_project/features/comment/repository"
	gameRepository "mini_project/features/game/repository"
	genreRepository "mini_project/features/genre/repository"
	recommendationRepository "mini_project/features/recommendation/repository"
	transactionRepository "mini_project/features/transaction/repository"
	userRepository "mini_project/features/user/repository"
	wishlistRepository "mini_project/features/wishlist/repository"

	"gorm.io/gorm"
)

func InitMigrationMysql(db *gorm.DB) {
	db.AutoMigrate(
		&userRepository.User{},
		&adminRepository.Admin{},
		&genreRepository.Genre{},
		&gameRepository.Game{},
		&wishlistRepository.Wishlist{},
		&transactionRepository.Transaction{},
		&commentRepository.Comment{},
		&recommendationRepository.Recommendation{},
	)
}
