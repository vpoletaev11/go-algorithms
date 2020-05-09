package sort

// BubbleSort - algorithm that repeatedly steps through the list, compares adjacent elements and swaps them if they are in the wrong order.
func BubbleSort(slice []int) []int {
	sliceLen := len(slice)
	for 0 < sliceLen {
		sliceLen--

		for i := 0; i < sliceLen; i++ {
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}
	}

	return slice
}

// BubbleSortReverse - algorithm that repeatedly steps through the list, compares adjacent elements and swaps them if they are in the wrong order.
// Slice is sorting from larger to smaller.
func BubbleSortReverse(slice []int) []int {
	sliceLen := len(slice)
	for 0 < sliceLen {
		sliceLen--

		for i := 0; i < sliceLen; i++ {
			if slice[i] < slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}
	}

	return slice
}
