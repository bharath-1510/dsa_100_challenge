package day07

import (
	"reflect"
	"sort"
	"testing"
)

func TestMaximumSumSubArray(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 3, -8, 7, -1, 2, 3}, 11},
		{[]int{-2, -4}, -2},
		{[]int{5, 4, 1, 7, 8}, 25},
	}

	for _, test := range tests {
		result := MaximumSumSubArray(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestMaximumSumSubArrayWithIndex(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{2, 3, -8, 7, -1, 2, 3}, []int{11,3,6}},
		{[]int{-2, -4}, []int{-2,0,0}},
		{[]int{5, 4, 1, 7, 8}, []int{25,0,4}},
	}

	for _, test := range tests {
		result := MaximumSumSubArrayWithIndex(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestTwoSumUsingSortingWithTwoPointer(t *testing.T) {
	tests := []struct {
		input    []int
		target int
		expected []int
	}{
		{[]int{0, -1, 2, -3, 1},-2, []int{-3,1}},
		{[]int{-2,0,1},1, []int{0,1}},
		{[]int{5, 4, 1, 7, 8},5, []int{4,1}},
	}

	for _, test := range tests {
		result := TwoSumUsingSortingWithTwoPointer(test.input,test.target)
		if !checkArray(result, test.expected) {
			t.Errorf("For input %v & target %v, expected %v, but got %v", test.input,test.target, test.expected, result)
		}
	}
}

func TestTwoSumUsingHash(t *testing.T) {

	tests := []struct {
		input    []int
		target int
		expected []int
	}{
		{[]int{0, -1, 2, -3, 1},-2, []int{-3,1}},
		{[]int{-2,0,1},1, []int{0,1}},
		{[]int{5, 4, 1, 7, 8},5, []int{4,1}},
	}

	for _, test := range tests {
		result := TwoSumUsingHash(test.input,test.target)
		if !checkArray(result, test.expected) {
			t.Errorf("For input %v & target %v, expected %v, but got %v", test.input,test.target, test.expected, result)
		}
	}
}


func TestTwoSumBruteForce(t *testing.T) {
	tests := []struct {
		input    []int
		target int
		expected []int
	}{
		{[]int{0, -1, 2, -3, 1},-2, []int{-3,1}},
		{[]int{-2,0,1},1, []int{0,1}},
		{[]int{5, 4, 1, 7, 8},5, []int{4,1}},
	}

	for _, test := range tests {
		result := TwoSumBruteForce(test.input,test.target)
		if !checkArray(result, test.expected) {
			t.Errorf("For input %v & target %v, expected %v, but got %v", test.input,test.target, test.expected, result)
		}
	}
}

func TestTripletSumBruteForce(t *testing.T) {
	tests := []struct {
		input    []int
		target int
		expected [][]int
	}{
		{[]int{0, -1, 2, -3, 1,3},-2, [][]int{{0,-3,1},{-1,2,-3}}},
		{[]int{-2,0,1,4,5},4, [][]int{{-2,1,5}}},
		{[]int{5, 4, 1, 7, -8},0, [][]int{{1,7,-8}}},
	}

	for _, test := range tests {
		result := TripletSumBruteForce(test.input,test.target)
		if !checkArrayTriplet(result, test.expected) {
			t.Errorf("For input %v & target %v, expected %v, but got %v", test.input,test.target, test.expected, result)
		}
	}
}

func TestTripletSumUsingHashing(t *testing.T) {
	tests := []struct {
		input    []int
		target int
		expected [][]int
	}{
		{[]int{0, -1, 2, -3, 1,3},-2, [][]int{{0,-3,1},{-1,2,-3}}},
		{[]int{-2,0,1,4,5},4, [][]int{{-2,1,5}}},
		{[]int{5, 4, 1, 7, -8},0, [][]int{{1,7,-8}}},
	}

	for _, test := range tests {
		result := TripletSumUsingHashing(test.input,test.target)
		if !checkArrayTriplet(result, test.expected) {
			t.Errorf("For input %v & target %v, expected %v, but got %v", test.input,test.target, test.expected, result)
		}
	}
}

func TestTripletSumUsingSorting(t *testing.T) {
	tests := []struct {
		input    []int
		target int
		expected [][]int
	}{
		{[]int{0, -1, 2, -3, 1,3},-2, [][]int{{0,-3,1},{-1,2,-3}}},
		{[]int{-2,0,1,4,5},4, [][]int{{-2,1,5}}},
		{[]int{5, 4, 1, 7, -8},0, [][]int{{1,7,-8}}},
	}

	for _, test := range tests {
		result := TripletSumUsingSorting(test.input,test.target)
		if !checkArrayTriplet(result, test.expected) {
			t.Errorf("For input %v & target %v, expected %v, but got %v", test.input,test.target, test.expected, result)
		}
	}
}

func checkArray(arr1 []int , arr2 []int) bool{
	sort.Ints(arr1)
	sort.Ints(arr2)
	return reflect.DeepEqual(arr1,arr2)
}

func checkArrayTriplet(arr1 [][]int , arr2 [][]int) bool{

	for i,_ := range arr1 {
		sort.Ints(arr1[i])
	}
	for i,_ := range arr2 {
		sort.Ints(arr2[i])
	}

	compare := func(a, b []int) bool {
		for i := range a {
			if a[i] != b[i] {
				return a[i] < b[i]
			}
		}
		return false
	}

	sort.Slice(arr1, func(i, j int) bool {
		return compare(arr1[i], arr1[j])
	})

	sort.Slice(arr2, func(i, j int) bool {
		return compare(arr2[i], arr2[j])
	})
	return reflect.DeepEqual(arr1,arr2)
}