package router

import (
	urlcontroller "github.com/codeshaine/url-shortner/internal/controller"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/{url}", urlcontroller.HanldeRedirect)
	r.Post("/shorten", urlcontroller.HandleShorten)
	return r

}
