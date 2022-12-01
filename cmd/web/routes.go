package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/scottwcode/go-bnb-bookings/pkg/config"
	"github.com/scottwcode/go-bnb-bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// ** muxer from https://github.com/bmizerany/pat
	// which is a Sinatra style pattern muxer for Go's net/http library
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	// ** muxer from https://github.com/go-chi/chi
	// chi is a lightweight, idiomatic and composable router for building
	// Go HTTP services for connecting to various middleware handlers
	mux := chi.NewRouter()

	// ** Use Recoverer middleware handler
	//    https://pkg.go.dev/github.com/go-chi/chi/middleware#Recoverer
	mux.Use(middleware.Recoverer)
	// ** Use our own middleware to write to console
	// mux.Use(WriteToConsole)
	// ** Use NoSurf middleware https://github.com/justinas/nosurf
	// which  is an HTTP package for Go that helps you prevent
	// Cross-Site Request Forgery attacks.
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
