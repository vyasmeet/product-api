//	Package classification for Product API
//
//	Documentation for Product API
//
//     Schemes: http
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//	swagger:meta
package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vyasmeet/product-api/data"
)

type KeyProduct struct{}

// A list of products returns in the response
//	swagger:response productsResponse
type productsResponseWrapper struct {
	//	All products in the system
	//	in: body
	Body []data.Product
}

// Products is http handler
type Products struct {
	log *log.Logger
	v   *data.Validation
}

func NewProducts(log *log.Logger, v *data.Validation) *Products {
	return &Products{log, v}
}

var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

type GenericError struct {
	Message string `json:"message"`
}

type ValidationError struct {
	Messages []string `json:"messages"`
}

func getProductID(r *http.Request) int {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	return id
}
