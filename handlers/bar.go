package handlers

import (
	"net/http"
)

type BarHandler struct {
}

func NewBarHandler() *BarHandler {
	return &BarHandler{}
}

func (h *BarHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(200)
	rw.Write([]byte(`{"result":"ok"}`))
}
