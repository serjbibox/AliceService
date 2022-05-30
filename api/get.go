package api

import (
	"AliceService/devices"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	sr := UnlinkResponse{"Hello!", "message"}
	SendHttp(w, &sr)
}

func DevicesList(w http.ResponseWriter, r *http.Request) {
	log.Println("ID запроса: ", r.Header.Get("X-Request-Id"))
	d := devices.Device{
		ID:   "1",
		Name: "Humidity",
	}
	devicesList := make([]devices.Device, 0)
	devicesList = append(devicesList, d)
	dl := DevicesInfoResponse{
		RequestID: r.Header.Get("X-Request-Id"),

		PL: devices.PayLoad{
			UserID:  "001",
			Devices: devicesList,
		},
	}
	SendHttp(w, &dl)
}
