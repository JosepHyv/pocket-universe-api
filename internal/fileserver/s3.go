package fileserver

import (
	"context"
	"time"
	"fmt"
	"log"
	"pocket-universe/internal/config"
	"pocket-universe/internal/models"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type S3Engine struct {
	client *minio.Client
	bucketName string
}


func NewClient(cfg config.Config) ( *S3Engine, error) {
	engine := &S3Engine{
		bucketName: cfg.StorageBucket,
	}
	creds := credentials.NewStaticV4(cfg.StorageAccessKey, cfg.StorageSecretKey, "")

	var err error
	engine.client, err = minio.New(cfg.StorageHost, &minio.Options{
		Creds: creds,
		Secure: cfg.StorageUseSSL,
	})

	if err != nil {
		log.Print(err)
		return nil, fmt.Errorf("Al iniciar conexion con el servidor de archivos: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	err = engine.initBucket(ctx)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	return engine, nil
}

func (fs *S3Engine) initBucket(ctx context.Context) error {

	exists, err := fs.client.BucketExists(ctx, fs.bucketName)
	if err != nil && exists {
		return nil
	}

	err = fs.client.MakeBucket(ctx, fs.bucketName, minio.MakeBucketOptions{})
	if err != nil {
		return fmt.Errorf("Error inicializando bucket: %w", err)
	}
	return nil
}


func (fs *S3Engine) SaveImage() (models.ImageUrl, error){
	return "", nil
}

func (fs *S3Engine) DeleteImage(imageId string ) error {
	return nil
}
