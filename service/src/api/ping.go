package api

import (
	"net/http"
)

func Ping(w http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}
