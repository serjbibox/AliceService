package api

import (
	"AliceService/devices"
	"encoding/json"
	"net/http"
)

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

type DevicesInfoResponse struct {
	RequestID string          `json:"request_id" example:"123"`
	PL        devices.PayLoad `json:"payload"`
}

func (s *DevicesInfoResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&s)
}
