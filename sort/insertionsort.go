package sort

// InsertionSort - algorithm that divides an slice into two parts: sorted and unsorted.
// Initially sorted part - its first element in slice.
// Then algorithm consistently insert first element from unsorted part in sorted part without breaking sorted order.
func InsertionSort(slice []int) []int {
	for i := 1; i < len(slice); i++ {
		k := i
		for k > 0 && slice[k-1] > slice[k] {
			slice[k-1], slice[k] = slice[k], slice[k-1]
			k--
		}
	}
	return slice
}
