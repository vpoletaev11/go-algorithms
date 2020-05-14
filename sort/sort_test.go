package sort_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpoletaev11/go-algorithms/sort"
)

func TestBubbleSort(t *testing.T) {
	dataset := []int{10, 1, 3, 6, 5, 7, 2, 8, 9, 4}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sort.BubbleSort(dataset)
	assert.Equal(t, expected, dataset)
}

func TestHeapSort(t *testing.T) {
	dataset := []int{10, 1, 3, 6, 5, 7, 2, 8, 9, 4}
	expected := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	sort.HeapSort(dataset)
	assert.Equal(t, expected, dataset)
}

func TestInsertionSort(t *testing.T) {
	dataset := []int{10, 1, 3, 6, 5, 7, 2, 8, 9, 4}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sort.InsertionSort(dataset)
	assert.Equal(t, expected, dataset)
}

func TestQuickSort(t *testing.T) {
	dataset := []int{10, 1, 3, 6, 5, 7, 2, 8, 9, 4}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sort.QuickSort(dataset)
	assert.Equal(t, expected, dataset)
}
