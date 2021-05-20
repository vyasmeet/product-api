package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	log *log.Logger
}

func NewGoodbye(log *log.Logger) *Goodbye {
	return &Goodbye{log}
}

func (goodbye *Goodbye) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Write([]byte("Byeee"))
}
