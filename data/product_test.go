package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{}

	err := p.Validator()

	if err != nil {
		t.Fatal(err)
	}
}
