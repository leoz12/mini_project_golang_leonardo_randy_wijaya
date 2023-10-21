package migration

import (
	adminRepository "mini_project/features/admin/repository"
	userRepository "mini_project/features/user/repository"

	"gorm.io/gorm"
)

func InitMigrationMysql(db *gorm.DB) {
	db.AutoMigrate(&userRepository.User{})
	db.AutoMigrate(&adminRepository.Admin{})
	// auto migrate untuk table book
}
