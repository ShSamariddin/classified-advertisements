package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ShSamariddin/classified-advertisements/db"
	"github.com/ShSamariddin/classified-advertisements/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := db.NewPool()
	if err != nil {
		log.Fatalf("correctly initialized db pool is required to start: %v", err)
	}

	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Logger)

	// routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("WelCUM!"))
	})
	r.Post("/auth", func(w http.ResponseWriter, r *http.Request) {
		handler.Auth(ctx, w, r)
	})

	r.Route("/api/v1", func(r chi.Router) {
		// TOTO: Write here all REST
		r.Get("/ads", handler.GetAds(ctx))
		r.Get("/ads/{ad_id}", handler.GetAdById(ctx))
		r.Post("/ads", handler.AddAd(ctx))
		r.Put("/ads/{ad_id}", handler.UpdateAdd(ctx))
		r.Delete("/ads/{ad_id}", handler.DeleteAd(ctx))
	})
	_ = http.ListenAndServe("localhost:8080", r)
}
