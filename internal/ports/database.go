package ports

import (
	"context"
	"pocket-universe/internal/models"
)

type DatabaseInterface interface {
	CreateStar(ctx context.Context, star models.Star) error
	DeleteStar(ctx context.Context, starId string) error
}
