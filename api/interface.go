package api

import "net/http"

type ResponseInterface interface {
	Send(w http.ResponseWriter)
}
