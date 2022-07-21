package sorting

func BubbleSort(arr []int) []int {
	n := len(arr)
	for i, _ := range arr {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}

	}
	return arr
}
