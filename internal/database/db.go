package database

import (
	"log"

	"github.com/introxx/myhttp/config"
	"github.com/introxx/myhttp/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *config.Config) {
	var err error
	DB, err = gorm.Open(postgres.Open(cfg.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error: Gorm.Open: %v", err)
	}

	// Включаем расширение для UUID (если его нет)
	if err := DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		log.Fatalf("Error creating uuid-ossp extension: %v", err)
	}

	// 🚀 Автоматически создаем / обновляем таблицы
	err = DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.ProductTranslation{},
		&models.ProductImage{},
	)
	if err != nil {
		log.Fatalf("Error on migration: %v", err)
	}

	log.Println("✅ Database connected and all models migrated successfully")
}
