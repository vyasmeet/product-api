package handlers

import (
	"net/http"

	"github.com/vyasmeet/product-api/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

//	create handles POST requests to add new product
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.log.Printf("[DEBUG] Inserting product: %#v\n", prod)
	data.AddProduct(prod)
}
