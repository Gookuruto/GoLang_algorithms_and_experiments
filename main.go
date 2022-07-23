package main

import (
	"fmt"

	"github.com/Gookuruto/GoLang_algorithms_and_experiments/sorting"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("Staring main function.")
	arr := []int{64, 25, 12, 22, 11}
	sugar.Infow("Initialised array. ", "array", arr)
	sugar.Info("Initialised array ", arr)

	result := sorting.SelectionSort(arr)

	table := []int{0, 11, 2, 4, 3, 50, 41, 41, 11, 2, 2, 2, 41}

	sorting.MergeSort(arr)
	fmt.Println(arr)

	fmt.Println(table)

	bubble_result := sorting.BubbleSort(arr)
	fmt.Println(result)
	fmt.Println(bubble_result)
}
