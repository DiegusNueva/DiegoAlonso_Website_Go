package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DiegusNueva/DiegoAlonso_Website_Go/internal/config"
	"github.com/DiegusNueva/DiegoAlonso_Website_Go/internal/routes"
)

func main() {
	cfg := config.LoadConfig()

	routes.RegisterRoutes()

	addr := fmt.Sprintf(":%s", cfg.Port)

	server := &http.Server{
		Addr:    addr,
		Handler: nil,
	}

	fmt.Println("Servidor corriendo en http://localhost:" + cfg.Port)

	log.Fatal(server.ListenAndServe())
}
