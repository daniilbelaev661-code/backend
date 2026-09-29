package main

import (
	"log"
	"net/http"

	"backend/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	// базовые middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// CORS — чтобы фронт с :3000 мог ходить на :8080
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: false,
	}))

	r.Route("/api", func(r chi.Router) {
		r.Post("/save", handlers.SaveHandler)
		r.Get("/read", handlers.ReadHandler) // <-- новое в v2.0
	})

	log.Println("Backend 2.0 (chi) running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}