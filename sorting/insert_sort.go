package sorting

func InsertSort(arr []int) []int {
	for i, v := range arr {
		key := v   //current value from array to be inserted to ordered part
		j := i - 1 // ordered part last index
		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j] // if inserted value is lower than currently selected item in ordered array move this item up by one
			j -= 1            // select smaler number in ordered array
		}
		arr[j+1] = key // insert key to it's place inordered array
	}
	return arr
}
