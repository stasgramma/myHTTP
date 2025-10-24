package storage

import (
	"context"
	"log"

	"github.com/introxx/myhttp/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func InitMinio(cfg *config.Config) {
	client, err := minio.New(cfg.MinIO.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinIO.AccessKey, cfg.MinIO.SecretKey, ""),
		Secure: cfg.MinIO.UseSSL,
	})
	if err != nil {
		log.Fatalf("Error initializing MinIO: %v", err)
	}

	MinioClient = client
	log.Println("Connected to MinIO successfully")

	ctx := context.Background()

	// Проверяем, существует ли bucket
	exists, err := client.BucketExists(ctx, cfg.MinIO.Bucket)
	if err != nil {
		log.Fatalf("Error checking MinIO bucket: %v", err)
	}

	if !exists {
		if err := client.MakeBucket(ctx, cfg.MinIO.Bucket, minio.MakeBucketOptions{}); err != nil {
			log.Fatalf("Error creating MinIO bucket: %v", err)
		}
		log.Printf("🪣 Created bucket: %s\n", cfg.MinIO.Bucket)
	} else {
		log.Printf("🪣 Bucket '%s' already exists\n", cfg.MinIO.Bucket)
	}
}
