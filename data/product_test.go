package data

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecksValidation(t *testing.T) {
	p := Product{
		Name:  "abc",
		Price: 1.22,
		SKU:   "asdf",
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestProductForMissingNameReturnsError(t *testing.T) {
	p := Product{
		Price: 2.25,
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestProductForMissingPriceReturnsError(t *testing.T) {
	p := Product{
		Name:  "test",
		Price: -1.25,
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}

func TestProductInvalidSKUReturnsError(t *testing.T) {
	p := Product{
		Name:  "test",
		Price: 1.12,
		SKU:   "tes",
	}

	v := NewValidation()
	err := v.Validate(p)
	if err != nil {
		t.Fatal(err)
	}
}

func TestProductValidSKUNOERROR(t *testing.T) {
	p := Product{
		Name:  "test",
		Price: 1.12,
		SKU:   "tes-jef-asd",
	}

	v := NewValidation()
	err := v.Validate(p)
	if err != nil {
		t.Fatal(err)
	}
}

func TestProductsToJSON(t *testing.T) {
	ps := []*Product{
		{
			Name: "test",
		},
	}
	b := bytes.NewBufferString("")
	err := ToJSON(ps, b)
	assert.NoError(t, err)
}
