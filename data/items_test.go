package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {
	assert := assert.New(t)
	it := &Item{Name: "Chair", Price: 30.0, SKU: "abc-def-ghi"}

	err := it.Validate()
	assert.Nil(err, err)
}
