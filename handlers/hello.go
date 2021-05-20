package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	log *log.Logger
}

func NewHello(log *log.Logger) *Hello {
	return &Hello{log}
}

func (h *Hello) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	h.log.Println("Hello World")
	data, error := ioutil.ReadAll(request.Body)
	if error != nil {
		http.Error(responseWriter, "Ooops!!", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(responseWriter, "Hello %s\n", data)
}
