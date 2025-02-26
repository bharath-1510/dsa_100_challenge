package day06

import (
	"reflect"
	"testing"
)

func TestFindElementsNBy3Times(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{2, 2, 3, 1, 3, 2, 1, 1}, []int{1, 2}},
		{[]int{-5, 3, -5}, []int{-5}},
		{[]int{3, 2, 2, 4, 1, 4}, []int{}},
	}

	for _, test := range tests {
		result := FindElementsNBy3Times(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestFindElementsNBy3TimesHash(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{2, 2, 3, 1, 3, 2, 1, 1}, []int{1, 2}},
		{[]int{-5, 3, -5}, []int{-5}},
		{[]int{3, 2, 2, 4, 1, 4}, []int{}},
	}

	for _, test := range tests {
		result := FindElementsNBy3TimesHash(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestFindElementsNByKTimesHash(t *testing.T) {
	tests := []struct {
		input    []int
		k int
		expected []int
	}{
		{[]int{2, 2, 3, 1, 3, 2, 1, 1},3, []int{1, 2}},
		{[]int{-5, 3, -5, 3},2, []int{-5,3}},
		{[]int{3, 2, 2, 4, 1, 4,4,4},2, []int{4}},
	}

	for _, test := range tests {
		result := FindElementsNByKTimesHash(test.input, test.k)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v , k= %v, expected %v, but got %v", test.input, test.k, test.expected, result)
		}
	}
}

func TestMostFrequentElementArrayBruteForce(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 2, 3, 1, 3, 2, 1, 1},1},
		{[]int{-5, 3, -1, 3,1},3},
		{[]int{3, 2, 2, 4, 1, 4, 4, 4},4},
	}

	for _, test := range tests {
		result := MostFrequentElementArrayBruteForce(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestMostFrequentElementArray(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 2, 3, 1, 3, 2, 1, 1},1},
		{[]int{-5, 3, -1, 3,1},3},
		{[]int{3, 2, 2, 4, 1, 4, 4, 4},4},
	}

	for _, test := range tests {
		result := MostFrequentElementArray(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestMostFrequentElementArrayUsingSorting(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 2, 3, 1, 3, 2, 1, 1},1},
		{[]int{-5, 3, -1, 3,1},3},
		{[]int{3, 2, 2, 4, 1, 4, 4, 4},4},
	}

	for _, test := range tests {
		result := MostFrequentElementArrayUsingSorting(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}
