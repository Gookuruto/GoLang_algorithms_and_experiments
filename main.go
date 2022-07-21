package main

import (
	"algorithm/sorting"
	"fmt"
)

func main() {
	arr := []int{64, 25, 12, 22, 11}

	result := sorting.SelectionSort(arr)

	bubble_result := sorting.BubbleSort(arr)
	fmt.Println(result)
	fmt.Println(bubble_result)
}
