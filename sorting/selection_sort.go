package sorting

func SelectionSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		min_index := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min_index] {
				min_index = j
			}
		}
		arr[min_index], arr[i] = arr[i], arr[min_index]
	}
	return arr

}
