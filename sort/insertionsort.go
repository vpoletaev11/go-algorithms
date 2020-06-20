package sort

/*
	INSERTION SORT:

	Best:    O(n)
	Average: O(n^2)
	Worst:   O(n^2)

	Memory:  O(1)
	Stable:  Yes


	goos: linux
	goarch: amd64
	BenchmarkInsertionSort-8        12972621                87.7 ns/op             0 B/op          0 allocs/op
*/

// InsertionSort - algorithm that divides an slice into two parts: sorted and unsorted.
// Initially sorted part - its first element in slice.
// Then algorithm consistently insert first element from unsorted part in sorted part without breaking sorted order.
func InsertionSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		k := i
		for k > 0 && slice[k-1] > slice[k] {
			slice[k-1], slice[k] = slice[k], slice[k-1]
			k--
		}
	}
}
