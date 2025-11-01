package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Price      float64        `gorm:"not null" json:"price"` // MDL
	Stock      int            `gorm:"not null" json:"stock"`
	IsHidden   bool           `gorm:"default:false" json:"is_hidden"`
	CategoryID *uint          `json:"category_id"` // связь с категорией (опционально)
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	// 🔗 Связи
	Translations []ProductTranslation `gorm:"constraint:OnDelete:CASCADE;" json:"translations"`
	Images       []ProductImage       `gorm:"constraint:OnDelete:CASCADE;" json:"images"`
}
