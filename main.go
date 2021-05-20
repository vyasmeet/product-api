package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/vyasmeet/product-api/handlers"
)

func main() {

	log := log.New(os.Stdout, "vyasmeet/product-api/", log.LstdFlags)

	productHandler := handlers.NewProducts(log)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", productHandler)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		log.Println("Starting server on port 9090")
		error := server.ListenAndServe()
		if error != nil {
			log.Fatal(error)
		}
	}()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	sig := <-signalChannel
	log.Println("Received terminate, graceful shutdown", sig)

	timeContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeContext)

}
