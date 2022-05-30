package api

import "net/http"

func EndpointPing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
