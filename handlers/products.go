package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/vyasmeet/product-api/data"
)

type Products struct {
	log *log.Logger
}

func NewProducts(log *log.Logger) *Products {
	return &Products{log}
}

func (product *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	// Get the list of products (GET)
	if r.Method == http.MethodGet {
		product.getProducts(rw, r)
		return
	}

	// Handle add product (POST)
	if r.Method == http.MethodPost {
		product.addProduct(rw, r)
		return
	}
	// Handle an update (PUT)
	if r.Method == http.MethodPut {

		regex := regexp.MustCompile(`/([0-9]+)`)
		group := regex.FindAllStringSubmatch(r.URL.Path, -1)

		if len(group) != 1 {
			product.log.Println("Invalid URI - More than one ID passed")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(group[0]) != 2 {
			product.log.Println("Invalid URI - More than one groups are captured")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := group[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			product.log.Println("Invalid URL - Cannot convert to number - ", idString)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		product.updateProducts(id, rw, r)
	}

	// CATCH OTHERS HERE - other methods not implemented yet
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

// handle GET request - Returns list of products
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle GET products")
	productList := data.GetProducts()
	error := productList.ToJSON(rw)
	if error != nil {
		http.Error(rw, "Unable to pasre product JSON", http.StatusInternalServerError)
		return
	}
}

// handle POST request - Add new product
func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle POST product")

	product := &data.Product{}
	error := product.FromJSON(r.Body)

	if error != nil {
		http.Error(rw, "Unable to create product from passed JSON", http.StatusBadRequest)
		return
	}

	data.AddProduct(product)
}

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.log.Println("Handle PUT product")

	product := &data.Product{}
	error := product.FromJSON(r.Body)

	if error != nil {
		http.Error(rw, "Unable to create product from passed JSON", http.StatusBadRequest)
		return
	}

	error = data.UpdateProduct(id, product)
	if error == data.ErrProductNotFound {
		http.Error(rw, "Product Not Found", http.StatusNotFound)
		return
	}

	if error != nil {
		http.Error(rw, "Product Not Found", http.StatusInternalServerError)
		return
	}
}
