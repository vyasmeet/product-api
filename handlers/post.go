package handlers

import (
	"net/http"

	"github.com/vyasmeet/product-api/data"
)

func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.log.Printf("[DEBUG] Inserting product: %#v\n", prod)
	data.AddProduct(prod)
}
