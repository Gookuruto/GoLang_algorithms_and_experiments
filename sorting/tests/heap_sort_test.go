package tests

import (
	"reflect"
	"testing"

	"workspace/sorting"
)

func TestHeapSort(t *testing.T) {
	table := []int{0, 11, 2, 4, 3, 50, 41}
	expected_result := []int{0, 2, 3, 4, 11, 41, 50}

	sorting.HeapSort(table)

	if !reflect.DeepEqual(table, expected_result) {
		t.Errorf("Bubblesort was incorrect, got: %#v, want: %#v.", table, expected_result)
	}

}

func TestIterHeapSort(t *testing.T) {
	table := []int{0, 11, 2, 4, 3, 50, 41}
	expected_result := []int{0, 2, 3, 4, 11, 41, 50}

	sorting.IterHeapSort(table)

	if !reflect.DeepEqual(table, expected_result) {
		t.Errorf("Bubblesort was incorrect, got: %#v, want: %#v.", table, expected_result)
	}

}
