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
	Message   string `json:"messadge" example:"123"`
}

func (s *UnlinkResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&s)
}

type Parameter struct {
	Instance string `json:"instance" example:"temperature"`
	Unit     string `json:"unit" example:"unit.temperature.celsius"`
}

type Property struct {
	Type        string    `json:"type" example:"123"`
	Retrievable bool      `json:"retrievable"`
	Reportable  bool      `json:"reportable"`
	Parameters  Parameter `json:"parameters" example:"123"`
}

type Capability struct {
	Type        string    `json:"type" example:"123"`
	Retrievable bool      `json:"retrievable"`
	Reportable  bool      `json:"reportable"`
	Parameters  Parameter `json:"parameters" example:"123"`
}

type DeviceInfo struct {
	Manufacturer string `json:"manufacturer" example:"Serj"`
	Model        string `json:"model" example:"S-01"`
	HwVersion    string `json:"hw_version" example:"V1.0"`
	SwVersion    string `json:"sw_version" example:"V1.1"`
}

type Device struct {
	ID           string            `json:"id" example:"123"`
	Name         string            `json:"name" example:"lamp"`
	Description  string            `json:"description" example:"123"`
	Room         string            `json:"room" example:"123"`
	Type         string            `json:"type" example:"devices.properties.float"`
	Custom_data  map[string]string `json:"custom_data"`
	Capabilities []Capability      `json:"Capabilities"`
	Properties   []Property        `json:"properties"`
}

type PayLoad struct {
	UserID  string   `json:"user_id" example:"123"`
	Devices []Device `json:"devices"`
}

type DevicesInfoResponse struct {
	RequestID string  `json:"request_id" example:"123"`
	PL        PayLoad `json:"payload"`
}

func (s *DevicesInfoResponse) Send(w http.ResponseWriter) {
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
		r.Get("/v1.0/user/devices", DevicesList)
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
	sr := UnlinkResponse{"Hello!", "message"}
	SendHttp(w, &sr)
}
func EndpointPing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func Unlink(w http.ResponseWriter, r *http.Request) {
	log.Println("ID запроса: ", r.Header.Get("X-Request-Id"))
	u := UnlinkResponse{
		RequestID: r.Header.Get("X-Request-Id"),
		Message:   r.Header.Get("Authorization"),
	}
	SendHttp(w, &u)
}
func DevicesList(w http.ResponseWriter, r *http.Request) {
	log.Println("ID запроса: ", r.Header.Get("X-Request-Id"))
	dl := DevicesInfoResponse{
		RequestID: r.Header.Get("X-Request-Id"),

		PL: PayLoad{
			UserID:  "001",
			Devices: []Device{},
		},
	}
	//r.Header.Get("X-Request-Id")
	//SendHttp(w, UnlinkResponse{"wtf"})
	SendHttp(w, &dl)
}
