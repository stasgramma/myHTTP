package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/introxx/myhttp/config"
	"github.com/introxx/myhttp/internal/models"
)

var DB *gorm.DB

func Connect(cfg *config.Config) {
	var err error
	DB, err = gorm.Open(postgres.Open(cfg.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error: Gorm.Open: %v", err)
	}

	// Авто-миграция таблиц
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Error on migration: %v", err)
	}
}
