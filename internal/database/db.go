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

	// Включаем расширение для генерации UUID в PostgreSQL
	if err := DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		log.Fatalf("Error creating uuid-ossp extension: %v", err)
	}

	// ❗ Удаляем таблицу, если она случайно существует с неправильным типом
	if DB.Migrator().HasTable(&models.User{}) {
		if err := DB.Migrator().DropTable(&models.User{}); err != nil {
			log.Fatalf("Error dropping users table: %v", err)
		}
		log.Println("Old users table dropped")
	}

	// Авто-миграция создаст таблицу заново с UUID
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Error on migration: %v", err)
	}

	log.Println("Database connected and migrated successfully")
}
