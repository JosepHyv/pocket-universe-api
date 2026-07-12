package http

import (
	"net/http"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "pocket-universe/docs"
)

type Api struct {}

func CreateServer() *Api {
	return &Api{}
}

func (api *Api) SetUpRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/docs/", httpSwagger.Handler(
	httpSwagger.URL("http://localhost:3000/docs/doc.json"),))
	mux.HandleFunc("GET /api/v1/health", api.HealthCheck)
	return mux

}
