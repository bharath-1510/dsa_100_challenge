package day01

import (
	"reflect"
	"testing"
)
// Rearrange the Postive and Negative Number using insertion sort
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

// Rearrange the Postive and Negative Number using Two Arrays
func TestRearrangeArrayUsingTwoArrays(t *testing.T) {
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
		rearrangeArray_using_two_arrays(arr)
		if !reflect.DeepEqual(arr, tc.want) {
			t.Errorf("rearrangeArray_using_two_arrays(%v) = %v; want %v", tc.input, arr, tc.want)
		}
	}
}

// Sort 0's ,1's and 2's using Dutch National Flag
func TestSort012UsingDNF(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{input: []int{1, 2, 0, 1, 2}, want: []int{0, 1, 1, 2, 2}},
		{input: []int{0, 1, 2, 0, 1, 2}, want: []int{0, 0, 1, 1, 2, 2}},
		{input: []int{0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1}, want: []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2}},
		{input: []int{}, want: []int{}},
	}

	for _, tc := range tests {
		arr := make([]int, len(tc.input))
		copy(arr, tc.input)
		sort_0_1_2_using_dnf(arr)
		if !reflect.DeepEqual(arr, tc.want) {
			t.Errorf("sort_0_1_2_using_dnf(%v) = %v; want %v", tc.input, arr, tc.want)
		}
	}
}

