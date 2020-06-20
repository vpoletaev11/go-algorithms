package search

/*
	LINEAR SEARCH:

	Best:    O(1)
	Average: O(n)
	Worst:   O(n)

	Memory:  O(1)


	goos: linux
	goarch: amd64
	BenchmarkLinearSearch-8                 275693473                4.30 ns/op            0 B/op          0 allocs/op
	BenchmarkLinearMultipleSearch-8         11093736               102 ns/op              56 B/op          3 allocs/op
*/

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
