package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Test",
		Price: 1.5,
		SKU:   "test-valid-sku-format",
	}

	err := p.Validator()

	if err != nil {
		t.Fatal(err)
	}
}
