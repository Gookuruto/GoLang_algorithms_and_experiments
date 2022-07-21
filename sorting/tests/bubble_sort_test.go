package tests

import (
	"reflect"
	"testing"

	"github.com/Gookuruto/GoLang_algorithms_and_experiments/sorting"
)

func TestBubbleSort(t *testing.T) {
	table := []int{0, 11, 2, 4, 3, 50, 41}
	expected_result := []int{0, 2, 3, 4, 11, 41, 50}

	result := sorting.BubbleSort(table)

	if !reflect.DeepEqual(result, expected_result) {
		t.Errorf("Bubblesort was incorrect, got: %#v, want: %#v.", result, expected_result)
	}

}
