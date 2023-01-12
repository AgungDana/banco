package orm

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	Id        uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	CraetedBy uint           `json:"craetedBy,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	UpdatedBy uint           `json:"updatedBy,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	DeletedBy uint           `json:"deletedBy,omitempty"`
}
