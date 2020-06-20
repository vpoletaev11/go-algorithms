package search

/*
	BINARY SEARCH:           BINARY MULTIPLE SEARCH:

	Best:    O(1)            Best:    O(1)
	Average: O(log n)        Average: O(log n + k)  where k - count of duplicate values
	Worst:   O(log n)        Worst:   O(log n + k)  where k - count of duplicate values

	Memory:  O(1)            Memory:  O(1)


	goos: linux
	goarch: amd64
	BenchmarkBinarySearch-8                 224958303                5.21 ns/op            0 B/op          0 allocs/op
	BenchmarkBinaryMultipleSearch-8         189608084                6.26 ns/op            0 B/op          0 allocs/op
*/

// BinarySearch - algorithm that finds middle value of slice and compare it with desired value.
// If desired value == middle value - desired value are found.
// If desired value < middle value - desired value should be in left part of slice from middle value.
// If desired value > middle value - desired value should be in right part of slice from middle value.
// Algorithm works recursively for divided parts.
// ALGORITHM WORKS ONLY ON SORTED DATA.
func BinarySearch(slice []int, desired int) int {
	begin, end := 0, len(slice)-1
	var middle int

	for {
		middle = (begin + end) / 2

		switch {
		case slice[middle] == desired:
			return middle

		case desired < slice[middle]:
			end = middle - 1

		case desired > slice[middle]:
			begin = middle + 1

		}

		if begin > end {
			return -1
		}
	}
}

// BinaryMultipleSearch same as BinarySearch, but when desired value found - algorithm sequentially checks it left and right neighbors to find range of desired values.
func BinaryMultipleSearch(slice []int, desired int) (from, to int) {
	begin, end := 0, len(slice)-1
	var middle int

	for {
		middle = (begin + end) / 2

		switch {
		case slice[middle] == desired:
			return fromToIndexes(slice, middle, desired)

		case desired < slice[middle]:
			end = middle - 1

		case desired > slice[middle]:
			begin = middle + 1

		}

		if begin > end {
			return -1, -1
		}
	}
}

// fromToIndexes finds range of desired values by checking of left and right neighbors of inputted desired value.
func fromToIndexes(slice []int, foundValIndex, desired int) (from, to int) {
	from = foundValIndex
	for {
		if from <= 0 {
			break
		}
		if slice[from-1] != desired {
			break
		}
		from--
	}

	to = foundValIndex
	for {
		if to >= len(slice)-1 {
			break
		}
		if slice[to+1] != desired {
			break
		}
		to++
	}
	return from, to
}
