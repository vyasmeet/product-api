package handlers

import (
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
		http.Error(rw, "Unable to pasre product JSON", http.StatusInternalServerError)
		return
	}
}

// handle POST request - Add new product
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle POST product")

	product := &data.Product{}
	error := product.FromJSON(r.Body)

	if error != nil {
		http.Error(rw, "Unable to create product from passed JSON", http.StatusBadRequest)
		return
	}

	data.AddProduct(product)
}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Unable to convert ID", http.StatusBadRequest)
		return
	}

	p.log.Println("Handle PUT product", id)

	product := &data.Product{}
	err = product.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to create product from passed JSON", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, product)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product Not Found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product Not Found", http.StatusInternalServerError)
		return
	}
}
