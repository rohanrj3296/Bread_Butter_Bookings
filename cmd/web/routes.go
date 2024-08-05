package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rohanrj3296/Bread_Butter_Bookings/pkg/config"
	"github.com/rohanrj3296/Bread_Butter_Bookings/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	//creating a file server
	fileServer := http.FileServer(http.Dir("./static/"))
	//starting the file server
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
