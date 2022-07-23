package sorting

func buildMaxHeap(arr []int) {
	for i, v := range arr {
		if v > arr[int((i-1)/2)] {
			j := i
			for arr[j] > arr[int((j-1)/2)] {
				arr[j], arr[int((j-1)/2)] = arr[int((j-1)/2)], arr[j]
				j = int((j - 1) / 2)
			}
		}
	}
}

func IterHeapSort(arr []int) {
	buildMaxHeap(arr)
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]

		j, index := 0, 0
		for {
			index = 2*j + 1
			if index < (i-1) && arr[index] < arr[index+1] {
				index++
			}
			if index < i && arr[j] < arr[index] {
				arr[j], arr[index] = arr[index], arr[j]
			}
			j = index
			if index >= i {
				break
			}
		}
	}
}
