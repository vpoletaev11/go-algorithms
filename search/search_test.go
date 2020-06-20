package search_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-algorithms/search"
)

var datasetUnsorted = []int{9, 55, 55, 55, 8, 1, 9, 12, 88, 3}
var datasetSorted = []int{0, 9, 9, 9, 10, 14, 16, 23, 25, 30}

func TestLinearSearch(t *testing.T) {
	assert.Equal(t, 4, search.LinearSearch(datasetUnsorted, 8))     // success
	assert.Equal(t, -1, search.LinearSearch(datasetUnsorted, 1000)) // not found
}

func TestLinearMultipleSearch(t *testing.T) {
	assert.Equal(t, []int{0, 6}, search.LinearMultipleSearch(datasetUnsorted, 9)) // success
	assert.Equal(t, []int{}, search.LinearMultipleSearch(datasetUnsorted, 1000))  // not found
}

func TestBinarySearch(t *testing.T) {
	assert.Equal(t, 1, search.BinarySearch(datasetSorted, 9))     // success
	assert.Equal(t, -1, search.BinarySearch(datasetSorted, 1000)) // not found
}

func TestBinaryMultipleSearch(t *testing.T) {
	// success
	from, to := search.BinaryMultipleSearch(datasetSorted, 9)
	actual := []int{from, to}
	expected := []int{1, 3}
	assert.Equal(t, expected, actual)

	// not found
	from, to = search.BinaryMultipleSearch(datasetSorted, 1000)
	actual = []int{from, to}
	expected = []int{-1, -1}
	assert.Equal(t, expected, actual)
}

func TestBinaryMultipleSearchAvoidOutOfRange(t *testing.T) {
	dataset := []int{1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2}

	// left bound
	from, to := search.BinaryMultipleSearch(dataset, 1)
	actual := []int{from, to}
	expected := []int{0, 5}
	assert.Equal(t, expected, actual)

	// right bound
	from, to = search.BinaryMultipleSearch(dataset, 2)
	actual = []int{from, to}
	expected = []int{6, 10}
	assert.Equal(t, expected, actual)
}
