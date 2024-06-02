package main

import (
	"Weather-API-GoExpertPostGrad/internal/handler"
	"Weather-API-GoExpertPostGrad/internal/middleware"
	"fmt"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	r.Route("/{cep}", func(r chi.Router) {
		r.Use(middleware.CheckCepMiddleware)
		r.Get("/", handler.HandleGetTemperatureByCEP)
	})

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
