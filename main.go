package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type ResponseInterface interface {
	Send(w http.ResponseWriter)
}

func SendHttp(w http.ResponseWriter, v ResponseInterface) {
	v.Send(w)
}

// @Description Структура HTTP ответа метода GET /submitData/{id}/status
type StatusResponse struct {
	ID     string `json:"id" example:"123"`
	Status string `json:"status" example:"new"`
}

func (s StatusResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&s)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Get("/", Hello)
	httpPort := ":"
	//Чтение системной переменной PORT для деплоя на Heroku
	if env, ok := os.LookupEnv("PORT"); !ok {
		httpPort += "8080"
	} else {
		httpPort += env
	}
	log.Panic(http.ListenAndServe(httpPort, r))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	sr := StatusResponse{"1", "Hello!"}
	SendHttp(w, sr)
}
