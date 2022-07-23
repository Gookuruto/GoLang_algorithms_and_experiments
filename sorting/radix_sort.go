package sorting

func RadixSort(arr []int) {

	_, max := findMinAndMax(arr)

	exp := 1

	for max/exp > 1 {
		countSort(arr, exp)
		exp = exp * 10
	}

}

func countSort(arr []int, exp int) {

	output := make([]int, len(arr))
	count := make([]int, 10)

	for i, _ := range arr {
		index := arr[i] / exp
		count[index%10]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := len(arr) - 1; i >= 0; i-- {
		index := arr[i] / exp
		output[count[index%10]-1] = arr[i]
		count[index%10]--
	}
	for i, _ := range arr {
		arr[i] = output[i]
	}

}

func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
