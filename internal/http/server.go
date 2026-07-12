package http

import (
	"net/http"
)

type Api struct {}

func CreateServer() *Api {
	return &Api{}
}

func (api *Api) SetUpRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", api.HealthCheck)
	return mux

}
