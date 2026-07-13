package ports

type ImageUrl string

type FileServerInterface interface {
	SaveImage() ImageUrl
	DeleteImage(imageId string) error
}
