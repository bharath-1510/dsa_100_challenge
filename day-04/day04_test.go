// day04_test.go
package day04

import (
	"reflect"
	"testing"
)

// reverseArray_using_aux_array reverses the array by using an auxiliary slice.
func TestReverseArrayUsingAuxArray(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{input: []int{1, 2, 3, 4, 5}, want: []int{5, 4, 3, 2, 1}},
		{input: []int{10, 20, 30, 40}, want: []int{40, 30, 20, 10}},
		{input: []int{1}, want: []int{1}},
		{input: []int{}, want: []int{}},
	}

	for _, tc := range tests {
		// Create a copy so as not to modify the test case data.
		arr := make([]int, len(tc.input))
		copy(arr, tc.input)
		reverseArray_using_aux_array(arr)
		if !reflect.DeepEqual(arr, tc.want) {
			t.Errorf("reverseArray_using_aux_array(%v) = %v; want %v", tc.input, arr, tc.want)
		}
	}
}

// Test for the two-pointer approach.
// reverseArrayTwoPointer reverses the array in-place with constant extra space.
func TestReverseArrayTwoPointer(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{input: []int{1, 2, 3, 4, 5}, want: []int{5, 4, 3, 2, 1}},
		{input: []int{10, 20, 30, 40}, want: []int{40, 30, 20, 10}},
		{input: []int{1}, want: []int{1}},
		{input: []int{}, want: []int{}},
	}

	for _, tc := range tests {
		arr := make([]int, len(tc.input))
		copy(arr, tc.input)
		reverseArrayTwoPointer(arr)
		if !reflect.DeepEqual(arr, tc.want) {
			t.Errorf("reverseArrayTwoPointer(%v) = %v; want %v", tc.input, arr, tc.want)
		}
	}
}

func TestFindMaximumAndMinimum(t *testing.T) {
	tests := []struct {
		input []int
		min   int
		max   int
	}{
		{input: []int{3, 1, 5, 0, 4}, min: 0, max: 5},
		{input: []int{1, 2, 3, 4, 5}, min: 1, max: 5},
		{input: []int{10, 20, 30, 40}, min: 10, max: 40},
		{input: []int{1}, min: 1, max: 1},
		{input: []int{0}, min: 0, max: 0},
	}
	for _, tc := range tests {

		min, max := FindMaximumAndMinimum(tc.input)
		if min != tc.min || max != tc.max {
			t.Errorf("FindMaximumAndMinimum(%v) = (%d, %d); want (%d, %d)", tc.input, min, max, tc.min, tc.max)
		}
	}
}

func TestFindSecondSmallestAndSecondLargest(t *testing.T) {
	tests := []struct {
		input []int
		min2  int
		max2  int
	}{
		{input: []int{3, 1, 5, 0, 4}, min2: 1, max2: 4},
		{input: []int{1, 2, 3, 4, 5}, min2: 2, max2: 4},
		{input: []int{10, 20, 30, 40}, min2: 20, max2: 30},
		{input: []int{1}, min2: 1, max2: 1},
		{input: []int{0}, min2: 0, max2: 0},
	}
	for _, tc := range tests {
		min2, max2 := FindSecondSmallestAndSecondLargest(tc.input)
		if min2 != tc.min2 || max2 != tc.max2 {
			t.Errorf("FindSecondSmallestAndSecondLargest(%v) = (%d, %d); want (%d, %d)", tc.input, min2, max2, tc.min2, tc.max2)
		}
	}
}
