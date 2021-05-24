package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChecksValidation(t *testing.T) {
	p := Product{
		Name:  "abc",
		Price: 1.22,
	}

	v := NewValidation()
	err := v.Validate(p)
	assert.Len(t, err, 1)
}
