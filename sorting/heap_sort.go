package sorting

// To heapify subtree rooted at index i.
// n is size of heap

func heapify(arr []int, n, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2
	//check if children are bigger than parrent if it's true chnage the largest index to the index of larger child
	if l < n && arr[largest] < arr[l] {
		largest = l
	}
	if r < n && arr[largest] < arr[r] {
		largest = r
	}
	// when largest index was changed chnage values in array and run chech on next node recursivly
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}

}

func HeapSort(arr []int) {
	n := len(arr) //Get heap size

	//build max heap
	for i := n/2 - 1; i >= 0; i -= 1 {
		heapify(arr, n, i)
	}

	// One by one extract an element from heap
	//max elemnt from heap is moved to the end of array and removed from heap
	// Then new heap is checked if it's maximum
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}
