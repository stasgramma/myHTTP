package config

import (
	"fmt"
	"os"
)

// Config — общая структура конфигурации проекта
type Config struct {
	// === Database ===
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// === MinIO ===
	MinIO MinIOConfig

	// == Server ===
	Port string
}

// MinIOConfig — конфигурация MinIO клиента
type MinIOConfig struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	Bucket    string
	UseSSL    bool
}

// LoadConfig — загружает конфигурацию из переменных окружения
func LoadConfig() *Config {
	return &Config{
		// PostgreSQL
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "mydb"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),

		// MinIO
		MinIO: MinIOConfig{
			Endpoint:  getEnv("MINIO_ENDPOINT", "localhost:9000"),
			AccessKey: getEnv("MINIO_ACCESS_KEY", "minioadmin"),
			SecretKey: getEnv("MINIO_SECRET_KEY", "minioadmin"),
			Bucket:    getEnv("MINIO_BUCKET", "mybucket"),
			UseSSL:    getEnv("MINIO_USE_SSL", "false") == "true",
		},

		// Server
		Port: getEnv("PORT", "8080"),
	}
}

// GetDSN — формирует строку подключения к PostgreSQL
func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort, c.DBSSLMode,
	)
}

// getEnv — возвращает значение переменной окружения или дефолт
func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
