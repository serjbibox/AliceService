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
type UnlinkResponse struct {
	RequestID string `json:"request_id" example:"123"`
}

func (s UnlinkResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&s)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	//r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Route("/", func(r chi.Router) {
		r.Get("/", Hello)
		r.Head("/v1.0", EndpointPing)
		r.Post("/v1.0/user/unlink", Unlink)
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

func Hello(w http.ResponseWriter, r *http.Request) {
	sr := UnlinkResponse{"Hello!"}
	SendHttp(w, sr)
}
func EndpointPing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func Unlink(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Header.Get("X-Request-Id"))
	//r.Header.Get("X-Request-Id")
	SendHttp(w, UnlinkResponse{"wtf"})
	//SendHttp(w, UnlinkResponse{r.Header.Get("X-Request-Id")})
}
