package sort

/*
	QUICK SORT:

	Best:    O(n log n)
	Average: O(n log n)
	Worst:   O(n^2)

	Memory:  O(log n)
	Stable:  No


	goos: linux
	goarch: amd64
	BenchmarkQuickSort-8             6083380               218 ns/op               0 B/op          0 allocs/op
*/

// QuickSort - algorithm that finds pivot value and divides slice in two parts: lesser and greater than pivot.
// Algorithm works recursively for divided parts.
func QuickSort(slice []int) {
	qsort(slice, 0, len(slice)-1)
}

// qsort - recursive function that does quick sort.
func qsort(slice []int, begin, end int) {
	if begin < end {
		pivot := partition(slice, begin, end)
		qsort(slice, begin, pivot)
		qsort(slice, pivot+1, end)
	}
}

// partition - Hoare's partitioning.
func partition(slice []int, begin, end int) int {
	pivot := slice[(begin+end)/2]

	for {
		for {
			if slice[begin] >= pivot {
				break
			}
			begin++
		}

		for {
			if slice[end] <= pivot {
				break
			}
			end--
		}

		if begin >= end {
			return end
		}

		slice[begin], slice[end] = slice[end], slice[begin]
	}
}
