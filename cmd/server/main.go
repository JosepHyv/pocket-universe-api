package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"pocket-universe/internal/config"
	"pocket-universe/internal/database"
	"pocket-universe/internal/fileserver"
	server "pocket-universe/internal/http"
)

// @title           Pocket Universe Backend
// @version         1.0
// @description     Low-Dependency http server for pocket universe project written in Go.
// @host            localhost:3000
// @BasePath        /
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Fallo crítico cargando configuración: %v", err)
	}

	db, err := database.NewMongoClient(cfg)
	if err != nil {
		log.Fatalf("Error fatal levantando la base de datos: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer func() {
		log.Println("Cerrando conexiones del servidor...")
		if err := db.Disconnect(ctx); err != nil {
			log.Printf("Error durante la desconexion de Mongo: %v\n", err)
		} else {
			log.Println("Conexion a Mongo cerrada exitosamente.")
		}
	}()

	fs, err := fileserver.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error fatal conectando con el file storage: %v", err)
	}

	api := server.CreateServer(db, fs)
	mux := api.SetUpRoutes()

	puerto := fmt.Sprintf(":%d", cfg.AppPort) // %d es para enteros
	log.Printf("🚀 Arrancando el servidor de Pocket Universe en http://localhost%s\n", puerto)

	err = http.ListenAndServe(puerto, mux)

	if err != nil {
		log.Fatalf("Error deteniendo el servidor: %v", err)
	}
}
