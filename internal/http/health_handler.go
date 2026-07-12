package http

import (
	"net/http"
)

// HealthCheck godoc
// @Summary      check server status
// @Description  returns ok
// @Tags         system
// @Produce      plain
// @Success      200  {string}  string  "Server is running!"
// @Router       /api/v1/health [get]
func (api *Api) HealthCheck(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Server running!"))
}

