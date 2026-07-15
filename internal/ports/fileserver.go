package ports

import (
	"pocket-universe/internal/models"
)

type FileServerInterface interface {
	SaveImage() (models.ImageUrl, error)
	DeleteImage(imageId string) error
}
