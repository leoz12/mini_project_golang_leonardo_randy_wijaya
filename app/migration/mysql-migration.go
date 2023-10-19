package migration

import (
	"mini_project/features/user/repository"

	"gorm.io/gorm"
)

func InitMigrationMysql(db *gorm.DB) {
	db.AutoMigrate(&repository.User{})
	// auto migrate untuk table book
}
