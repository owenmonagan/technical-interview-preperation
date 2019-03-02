package quicksort

import (
	"fmt"
	"testing"
)

func TestQuicksort(t *testing.T) {
	data := []int{8, 10, 4, -1, 6, 123123, 5}
	fmt.Println(data)
	result := quicksort(data)
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

func TestQuicksortEdgeCase(t *testing.T) {
	data := []int{}
	fmt.Println(data)
	result := quicksort(data)
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
