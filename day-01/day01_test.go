package day01

import (
	"reflect"
	"testing"
)

func TestRearrangeArrayUsingInsertionSort(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{input: []int{1, 2, -3, 4, -5}, want: []int{-3, -5, 1, 2, 4}},
		{input: []int{12, 11, -13, -5, 6, -7, 5, -3, -6}, want: []int{-13, -5, -7, -3, -6, 12, 11, 6, 5}},
		{input: []int{1}, want: []int{1}},
		{input: []int{}, want: []int{}},
	}

	for _, tc := range tests {
		arr := make([]int, len(tc.input))
		copy(arr, tc.input)
		rearrangeArray_using_insertion_sort(arr)
		if !reflect.DeepEqual(arr, tc.want) {
			t.Errorf("rearrangeArray_using_insertion_sort(%v) = %v; want %v", tc.input, arr, tc.want)
		}
	}
}
