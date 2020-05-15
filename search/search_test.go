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

func TestBinarySearch(t *testing.T) {
	dataset := []int{1, 2, 4, 7, 9, 13, 15, 20, 23, 28, 31, 42, 48, 50}

	assert.Equal(t, 4, search.BinarySearch(dataset, 9))     // success
	assert.Equal(t, -1, search.BinarySearch(dataset, 1000)) // not found
}

func TestBinaryMultipleSearch(t *testing.T) {
	dataset := []int{1, 2, 3, 4, 5, 5, 5, 5, 6, 8, 15, 19, 20, 30, 48, 56, 100}

	// success
	from, to := search.BinaryMultipleSearch(dataset, 5)
	actual := []int{from, to}
	expected := []int{4, 7}
	assert.Equal(t, expected, actual)

	// not found
	from, to = search.BinaryMultipleSearch(dataset, 1000)
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

	// not found
	from, to = search.BinaryMultipleSearch(dataset, 2)
	actual = []int{from, to}
	expected = []int{6, 10}
	assert.Equal(t, expected, actual)
}
