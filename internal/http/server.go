package http

import (
	"net/http"
	"pocket-universe/internal/ports"

	_ "pocket-universe/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Api struct {
	database ports.DatabaseInterface
}

func CreateServer(db ports.DatabaseInterface) *Api {
	return &Api{
		database: db,
	}
}

func (api *Api) SetUpRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/docs/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/docs/doc.json"),
	))

	mux.HandleFunc("GET /api/v1/health", api.HealthCheck)
	mux.HandleFunc("POST /api/v1/createStar", api.CreateStar)

	return mux
}
