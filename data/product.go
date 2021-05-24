package data

import (
	"fmt"
)

var ErrProductNotFound = fmt.Errorf("Product Not Found")

// Product defines structure of API product
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this product
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU       string `json:"sku" validate:"sku"`
	CreatedOn string `json:"-"`
	UpdatedOn string `json:"-"`
	DeletedOn string `json:"-"`
}

//	slice/Array of Product
type Products []*Product

//	GetProducts returns all products from db
func GetProducts() Products {
	return productList
}

//	GetProductByID returns a single product which matches id from the db
//	If a product is not found, it returns ProductNotFound Error
func GetProductsByID(id int) (*Product, error) {
	i := findIndexByProductID(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}
	return productList[i], nil
}

//	AddProduct adds new product in db
func AddProduct(p *Product) {
	maxID := productList[len(productList)-1].ID
	p.ID = maxID + 1
	productList = append(productList, p)
}

//	UpdateProduct replaces a product in the db with given product.
//	If a product with given id does not exists in db, this returns ProductNotFound Error
func UpdateProduct(p *Product) error {
	i := findIndexByProductID(p.ID)
	if i == -1 {
		return ErrProductNotFound
	}
	productList[i] = p
	return nil
}

//	DeleteProduct deletes a product from db
func DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return ErrProductNotFound
	}
	productList = append(productList[:i], productList[i+1])
	return nil
}

//	findIndexByProductID finds the index of a product in the db
//	returns -1 when no product found for given ID
func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}
	return -1
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc123",
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "edc345",
	},
}
