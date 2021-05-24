package handlers

import (
	"net/http"

	"github.com/vyasmeet/product-api/data"
)

//	swagger:route GET /products products listProducts
//	Returns a list of products
//	responses:
//	200: productsResponse
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("[DEBUG] get all records")
	products := data.GetProducts()

	err := data.ToJSON(products, rw)
	if err != nil {
		p.log.Println("[ERROR] serialising product", err)
	}
}

func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.log.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductsByID(id)

	switch err {
	case nil:
	case data.ErrProductNotFound:
		p.log.Println("[ERROR] fetching product", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.log.Println("[ERROR] fetching product", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prod, rw)
	if err != nil {
		p.log.Println("[ERROR] serializing product", err)
	}
}
