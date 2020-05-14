package search_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-algorithms/search"
)

func TestLinearSearch(t *testing.T) {
	dataset := []int{4, 7, 1, 3, 11, 82, 54, 6, 6, 14, 9, 13, 28}

	assert.Equal(t, 10, search.LinearSearch(dataset, 9))    // success
	assert.Equal(t, -1, search.LinearSearch(dataset, 1000)) // not found
}

func TestLinearMultipleSearch(t *testing.T) {
	dataset := []int{4, 7, 1, 3, 11, 82, 54, 6, 6, 14, 9, 13, 28}

	assert.Equal(t, []int{7, 8}, search.LinearMultipleSearch(dataset, 6)) // success
	assert.Equal(t, []int{}, search.LinearMultipleSearch(dataset, 1000))  // not found
}
