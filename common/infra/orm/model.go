package orm

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	Id        uint `gorm:"primarykey"`
	CreatedAt time.Time
	CraetedBy uint
	UpdatedAt time.Time
	UpdatedBy uint
	DeletedAt gorm.DeletedAt `gorm:"index"`
	DeletedBy uint
}
