package database

import (
	"context"
	"fmt"
	"pocket-universe/internal/config"
	"pocket-universe/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
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
	engine := &MongoEngine{
		dbName: cfg.DatabaseName,
	}

	var err error

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseHost, cfg.DatabasePort)

	engine.client, err = mongo.Connect(options.Client().ApplyURI(uri))

	if err != nil {
		return nil, fmt.Errorf("realizando la conexion: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = engine.client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("conexion perdida o no establecida: %w", err)
	}

	return engine, nil
}

func (db *MongoEngine) Disconnect(ctx context.Context) error {
	if db.client == nil {
		return nil
	}
	return db.client.Disconnect(ctx)
}
