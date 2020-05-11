package sort

import "github.com/vpoletaev11/go-data-structures/tree"

// HeapSort - algorithm that consistently push elements from slice in the heap and write pop result in the same slice.
// In this realization values in sorted slice will be in order from max to min.
func HeapSort(slice []int) {
	heap := tree.Heap{}
	for _, val := range slice {
		heap.Push(val)
	}

	for i := 0; i < len(slice); i++ {
		slice[i], _ = heap.Pop()
	}
}
