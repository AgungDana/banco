package orm

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	Id        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UpdatedBy string
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
