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

	bubble_result := sorting.BubbleSort(arr)
	fmt.Println(result)
	fmt.Println(bubble_result)
}
