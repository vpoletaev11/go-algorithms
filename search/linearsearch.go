package search

// LinearSearch - algorithm that sequentially checks each element of the list until a match is found.
func LinearSearch(slice []int, desired int) int {
	for i, val := range slice {
		if val == desired {
			return i
		}
	}
	return -1
}

// LinearMultipleSearch - algorithm that sequentially checks each element of whole slice.
func LinearMultipleSearch(slice []int, desired int) []int {
	indexes := []int{}
	for i, val := range slice {
		if val == desired {
			indexes = append(indexes, i)
		}
	}
	return indexes
}
