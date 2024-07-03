package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/igorfarodrigues/api-postgres/configs"
	"github.com/igorfarodrigues/api-postgres/handlers"
)

func main() {
	err := configs.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Log the loaded configuration
	log.Printf("Loaded configuration: %+v", configs.GetDB())

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	port := configs.GetServerPort()
	log.Printf("Server starting on port %s", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		log.Fatal(err)
	}

}
