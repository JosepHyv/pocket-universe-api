package database

import (
	"context"
	"pocket-universe/internal/config"
	"pocket-universe/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoEngine struct {
	client *mongo.Client
	dbName string
}

func (db *MongoEngine) CreateStar(ctx context.Context, star models.Star) error {
	return nil
}

func (db *MongoEngine) DeleteStar(ctx context.Context, starId string) error {
	return nil
}

func NewMongoClient(cfg config.Config) (*MongoEngine, error) {

}
