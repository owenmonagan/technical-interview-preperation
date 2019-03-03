package mergesort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	data := []int{8, 10, 4, -1, -1, -1, 8, 6, 123123, 5}
	fmt.Println(data)
	result := Mergesort(data)
	fmt.Println(result)

	if len(data) != len(result) {
		t.Error("incorrect return length for quicksort")
	}
	for index := range data {
		if index+1 != len(data) {
			if result[index] > result[index+1] {
				t.Error("list not sorted")
			}
		}
	}
}

func TestSortEdgeCase(t *testing.T) {
	data := []int{}
	fmt.Println(data)
	result := Mergesort(data)
	fmt.Println(result)

	if len(data) != len(result) {
		t.Error("incorrect return length for quicksort")
	}
	for index := range data {
		if index+1 != len(data) {
			if result[index] > result[index+1] {
				t.Error("list not sorted")
			}
		}
	}
}
