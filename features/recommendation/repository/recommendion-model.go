package recommendationRepository

import (
	"time"

	"gorm.io/gorm"
)

type Recommendation struct {
	ID        string `gorm:"primarykey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
