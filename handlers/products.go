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
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vyasmeet/product-api/data"
)

type Products struct {
	log *log.Logger
}

func NewProducts(log *log.Logger) *Products {
	return &Products{log}
}

// handle GET request - Returns list of products
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle GET products")
	productList := data.GetProducts()
	error := productList.ToJSON(rw)
	if error != nil {
		p.log.Println("Unable to pasre product JSON", error)
		http.Error(rw, "Unable to pasre product JSON", http.StatusInternalServerError)
		return
	}
}

// handle POST request - Add new product
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle POST product")

	product := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&product)
}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		p.log.Println("Unable to convert ID", err)
		http.Error(rw, "Unable to convert ID", http.StatusBadRequest)
		return
	}

	p.log.Println("Handle PUT product", id)

	product := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &product)
	if err == data.ErrProductNotFound {
		p.log.Println("Product Not Found", err)
		http.Error(rw, "Product Not Found", http.StatusNotFound)
		return
	}

	if err != nil {
		p.log.Println("Product Not Found", err)
		http.Error(rw, "Product Not Found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		product := data.Product{}
		error := product.FromJSON(r.Body)

		if error != nil {
			p.log.Println("Unable to create product from passed JSON", error)
			http.Error(rw, "Unable to create product from passed JSON", http.StatusBadRequest)
			return
		}

		// validate the product
		error = product.Validator()
		if error != nil {
			p.log.Println("Unable to validate product", error)
			http.Error(
				rw,
				fmt.Sprintf("Unable to validate product: %s", error),
				http.StatusBadRequest,
			)
			return
		}

		context := context.WithValue(r.Context(), KeyProduct{}, product)
		req := r.WithContext(context)

		next.ServeHTTP(rw, req)
	})

}
