package sorting

func CountSort(arr []int) []int {

	output := make([]int, len(arr))
	count := make([]int, 256)

	for _, v := range arr {
		count[v]++
	}

	for i := 1; i < 256; i++ {
		count[i] += count[i-1]
	}
	for i := 0; i < len(arr); i++ {
		output[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}
	return output
}
