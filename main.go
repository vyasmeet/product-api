package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/vyasmeet/product-api/data"
	"github.com/vyasmeet/product-api/handlers"
)

func main() {

	log := log.New(os.Stdout, "vyasmeet/product-api/", log.LstdFlags)
	v := data.NewValidation()

	productHandler := handlers.NewProducts(log, v)

	serveMux := mux.NewRouter()

	// GET subRouter
	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", productHandler.ListAll)
	getRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.ListSingle)

	// POST subrouter
	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", productHandler.Create)
	postRouter.Use(productHandler.MiddlewareValidateProduct)

	// PUT subrouter
	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.Update)
	putRouter.Use(productHandler.MiddlewareValidateProduct)

	// DELETE subRouter
	deleteRouter := serveMux.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.Delete)

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
