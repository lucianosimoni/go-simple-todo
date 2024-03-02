package main

import (
	"fmt"
	"net/http"

	"luciano/simple-todo/internal/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/user", routes.UserRouter)

	fmt.Println("🚀 Listening on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
