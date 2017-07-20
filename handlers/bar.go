package handlers

import (
	"net/http"
	"encoding/json"
)

type BarHandler struct {
}

func NewBarHandler() *BarHandler {
	return &BarHandler{}
}

func (h *BarHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(200)

	type Response struct {
		Result string `json:"result"`
	}

	b, _ := json.Marshal(Response{Result:"ok"})
	rw.Write(b)
}
