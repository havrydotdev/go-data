package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	id, err := r.Insert(Test{
		Name:  "test6",
		Count: 6,
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, 6, id)
}
