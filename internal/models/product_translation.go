package models

type ProductTranslation struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	ProductID   uint   `gorm:"not null;index" json:"product_id"`
	Language    string `gorm:"type:varchar(2);not null" json:"language"` // en, ru, ro
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`

	// 🔗 Обратная связь
	Product Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE;" json:"-"`
}
