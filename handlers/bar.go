package handlers

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
)

type BarHandler struct {
}

func NewBarHandler() *BarHandler {
	return &BarHandler{}
}

type RequestBar struct {
	Result string `json:"result"`
}

func (h *BarHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	reqBodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("error in read request" + err.Error())
		rw.Write([]byte(err.Error()))
	}
	var reqBody *RequestBar
	err = json.Unmarshal(reqBodyBytes, &reqBody)

	rw.WriteHeader(200)
	rw.Write([]byte(reqBody.Result))
}
