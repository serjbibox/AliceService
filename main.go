package main

import (
	"AliceService/api"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	//r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Get("/", api.Hello)
	r.Route("/v1.0", func(r chi.Router) {
		r.Head("/", api.EndpointPing)
		r.Post("/user/unlink", api.Unlink)
		r.Post("/user/devices/query", api.Query)
		r.Get("/user/devices", api.DevicesList)
	})
	httpPort := ":"
	//Чтение системной переменной PORT для деплоя на Heroku
	if env, ok := os.LookupEnv("PORT"); !ok {
		httpPort += "8080"
	} else {
		httpPort += env
	}

	log.Panic(http.ListenAndServe(httpPort, r))
}
