package models

import (
	"time"
)

// 🗂 Category — основная таблица категорий
type Category struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// Связь с переводами
	Translations []CategoryTranslation `gorm:"constraint:OnDelete:CASCADE"`
	// Связь с продуктами (чтобы можно было прикрепить продукт к категории)
	Products []Product
}

// 🌍 CategoryTranslation — переводы для категорий
type CategoryTranslation struct {
	ID         uint   `gorm:"primaryKey"`
	CategoryID uint   `gorm:"index;not null"`
	Language   string `gorm:"type:varchar(2);not null"` // en, ru, ro
	Name       string `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	// Связь с Category
	Category Category `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE"`
}
