package sorting

func MergeSort(arr []int) {
	if len(arr) > 1 {
		mid := len(arr) / 2
		l := make([]int, mid)
		r := make([]int, len(arr)-mid)
		//IMPORTANT PART WE NEED TO CREATE COPY OF SLICES because if we make it by reference we will change l and r every time we move someting to arr.
		//Keep in mind it's not a python and array is passed by reference
		copy(l, arr[:mid])
		copy(r, arr[mid:])

		MergeSort(l)
		MergeSort(r)

		i, j, k := 0, 0, 0

		for i < len(l) && j < len(r) {
			if l[i] <= r[j] {
				arr[k] = l[i]
				i++
			} else {
				arr[k] = r[j]
				j++
			}
			k++
		}
		for i < len(l) {
			arr[k] = l[i]
			i++
			k++
		}

		for j < len(r) {
			arr[k] = r[j]
			j++
			k++
		}
	}

}
