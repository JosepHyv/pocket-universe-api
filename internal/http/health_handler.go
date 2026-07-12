package http

import (
	"net/http"
)

func (api *Api) HealthCheck(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Server running!"))
}

