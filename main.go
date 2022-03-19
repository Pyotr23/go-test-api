package main

import (
	"api/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Use(middleware.Recoverer)
	r.Route("/users", func(r chi.Router) {
		r.Get("/", handlers.GetUsers)
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	http.ListenAndServe(":3000", r)
}
