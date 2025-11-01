package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	Price    float64 `gorm:"not null" json:"price"` // цена в MDL
	Stock    int     `gorm:"not null" json:"stock"`
	IsHidden bool    `gorm:"default:false" json:"is_hidden"`

	// 🔗 Категория
	CategoryID *uint    `gorm:"index" json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID;constraint:OnDelete:SET NULL" json:"category"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// 🌍 Переводы и изображения
	Translations []ProductTranslation `gorm:"constraint:OnDelete:CASCADE;" json:"translations"`
	Images       []ProductImage       `gorm:"constraint:OnDelete:CASCADE;" json:"images"`
}
