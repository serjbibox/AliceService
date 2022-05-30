package api

import (
	"log"
	"net/http"
)

func Unlink(w http.ResponseWriter, r *http.Request) {
	log.Println("ID запроса: ", r.Header.Get("X-Request-Id"))
	u := UnlinkResponse{
		RequestID: r.Header.Get("X-Request-Id"),
		Message:   r.Header.Get("Authorization"),
	}
	SendHttp(w, &u)
}

func Query(w http.ResponseWriter, r *http.Request) {
	log.Println("ID запроса: ", r.Body)
	u := UnlinkResponse{
		RequestID: r.Header.Get("X-Request-Id"),
		Message:   r.Header.Get("Authorization"),
	}
	SendHttp(w, &u)
}

func Action(w http.ResponseWriter, r *http.Request) {
	log.Println("ID запроса: ", r.Body)
	u := UnlinkResponse{
		RequestID: r.Header.Get("X-Request-Id"),
		Message:   r.Header.Get("Authorization"),
	}
	SendHttp(w, &u)
}
