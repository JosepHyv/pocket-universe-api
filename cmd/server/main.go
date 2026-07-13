package main

import (
	"log"
	"fmt"
	"net/http"
	"pocket-universe/internal/config"
	server "pocket-universe/internal/http"
)

// @title           Pocket Universe Backend
// @version         1.0
// @description     Low-Dependency http server for pocket universe project written in Go.
// @host            localhost:3000
// @BasePath        /
func main(){
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Print(err)
		return
	}

	api := server.CreateServer(cfg)
	mux := api.SetUpRoutes()

	log.Printf("🚀 Arrancando el servidor en http://localhost%s\n", fmt.Sprintf(":%v",cfg.AppPort))

	err = http.ListenAndServe(fmt.Sprintf(":%v", cfg.AppPort), mux)

	if err != nil {
		log.Fatalf("Error starting server %v", err)
	}

	log.Printf("Server started at: localhost%s", cfg.AppPort)


}
