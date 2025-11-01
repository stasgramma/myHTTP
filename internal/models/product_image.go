package models

type ProductImage struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	ProductID    uint   `gorm:"not null;index" json:"product_id"`
	ImageURL     string `gorm:"not null" json:"image_url"`
	DisplayOrder int    `gorm:"default:0" json:"display_order"`

	// 🔗 Обратная связь
	Product Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;" json:"-"`
}
