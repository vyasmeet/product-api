package handlers

import (
	"net/http"

	"github.com/vyasmeet/product-api/data"
)

func (p *Products) Delete(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)
	p.log.Println("[DEBUG] deleting record id", id)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		p.log.Println("[ERROR] deleting record id does not exist", id)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.log.Println("[ERROR] deleting record", err)
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
