package sort

/*
	BUBBLE SORT:

	Best:    O(n)
	Average: O(n^2)
	Worst:   O(n^2)

	Memory:  O(1)
	Stable:  Yes


	goos: linux
	goarch: amd64
	BenchmarkBubbleSort-8           11839004               125 ns/op               0 B/op          0 allocs/op
*/

// BubbleSort - algorithm that repeatedly steps through the list, compares adjacent elements and swaps them if they are in the wrong order.
func BubbleSort(slice []int) {
	sliceLen := len(slice)
	for 0 < sliceLen {
		sliceLen--

		for i := 0; i < sliceLen; i++ {
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}
	}
}
