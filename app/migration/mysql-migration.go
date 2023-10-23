package migration

import (
	adminRepository "mini_project/features/admin/repository"
	gameRepository "mini_project/features/game/repository"
	genreRepository "mini_project/features/genre/repository"
	userRepository "mini_project/features/user/repository"

	"gorm.io/gorm"
)

func InitMigrationMysql(db *gorm.DB) {
	db.AutoMigrate(&userRepository.User{}, &adminRepository.Admin{}, &genreRepository.Genre{}, &gameRepository.Game{})
}
