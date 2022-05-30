package api

import "net/http"

func SendHttp(w http.ResponseWriter, v ResponseInterface) {
	v.Send(w)
}
