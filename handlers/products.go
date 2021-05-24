package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vyasmeet/product-api/data"
)

//	KeyProduct is a key used for the Product object in the context
type KeyProduct struct{}

//	Products is http handler for getting and updating products
type Products struct {
	log *log.Logger
	v   *data.Validation
}

//	NewProducts returns a new products handler with given logger
func NewProducts(log *log.Logger, v *data.Validation) *Products {
	return &Products{log, v}
}

//	ErrInvalidProductPath is an error message when product path is not valid
var ErrInvalidProductPath = fmt.Errorf("invalid Path, path should be /products/[id]")

//	GenericError is a generic error message returned by server
type GenericError struct {
	Message string `json:"message"`
}

//	ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

//	getProductID returns the product ID from the URL
//	panics if cannot convert the id into integer
//	this should never happen as the router ensures that this is a valid number
func getProductID(r *http.Request) int {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	return id
}
