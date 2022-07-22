package tests

import (
	"reflect"
	"testing"

	"github.com/Gookuruto/GoLang_algorithms_and_experiments/sorting"
)

func TestQuickSort(t *testing.T) {
	table := []int{0, 11, 2, 4, 3, 50, 41, 41, 11, 2, 2, 2, 41}
	expected_result := []int{0, 2, 2, 2, 2, 3, 4, 11, 11, 41, 41, 41, 50}

	result := sorting.QuickSort(table)

	if !reflect.DeepEqual(table, expected_result) {
		t.Errorf("Bubblesort was incorrect, got: %#v, want: %#v.", result, expected_result)
	}

}
