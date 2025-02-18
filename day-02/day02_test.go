package day01

import (
	"reflect"
	"testing"
	"fmt"
)
// Stock Buy and Sell Problem
func TestStockBuySell(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{input: []int{100, 180, 260, 310, 40, 535, 695}, want: 655},
		{input: []int{4, 2, 2, 2, 4}, want: 2},
		{input: []int{4, 2}, want: 0},
	}

	for _, tc := range tests {
		arr := make([]int, len(tc.input))
		copy(arr, tc.input)
		res := stockSellBuy(arr)
		if !reflect.DeepEqual(res, tc.want) {
			t.Errorf("stockSellBuy(%v) = %v; want %v", tc.input, res, tc.want)
		}
	}
}

// Helper function to compare two slices for equality
func equalSlices(a, b []int) bool {

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestFindFirstandLastOccurence(t *testing.T) {
	tests := []struct {
		input    []int
		target   int
		expected []int
	}{
		{[]int{2, 2, 3, 3, 4, 5, 5}, 2, []int{0, 1}},
		{[]int{2, 2, 3, 3, 4, 4, 5}, 2, []int{0, 1}},
		{[]int{2, 3, 3, 4, 4, 5, 5}, 3, []int{1, 2}},
		{[]int{2, 2, 3, 4, 4, 5, 5}, 4, []int{3, 4}},
		{[]int{2, 2, 3, 3, 4, 5, 5, 6, 6}, 5, []int{5, 6}},
	}

	for _, test := range tests {
		first, last := FindFirstandLastOccurence(test.input, test.target)
		result := []int{first, last}
		if !equalSlices(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}

}

func TestFindFirstandLastOccurenceUsingBinarySearch(t *testing.T) {
	tests := []struct {
		input    []int
		target   int
		expected []int
	}{
		{[]int{2, 2, 3, 3, 4, 5, 5}, 2, []int{0, 1}},
		{[]int{2, 2, 3, 3, 4, 4, 5}, 2, []int{0, 1}},
		{[]int{2, 3, 3, 4, 4, 5, 5}, 3, []int{1, 2}},
		{[]int{2, 2, 3, 4, 4, 5, 5}, 4, []int{3, 4}},
		{[]int{2, 2, 3, 3, 4, 5, 5, 6, 6}, 5, []int{5, 6}},
	}

	for _, test := range tests {
		first, last := FindFirstandLastOccurenceUsingBinarySearch(test.input, test.target)
		result := []int{first, last}
		if !equalSlices(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}

}
