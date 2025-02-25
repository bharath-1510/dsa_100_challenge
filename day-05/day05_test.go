package day05

import (
	"reflect"
	"testing"
)

func TestMoveZeroEnd(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0, 0, 1}, []int{1, 0, 0}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{1, 0, 2, 0, 3}, []int{1, 2, 3, 0, 0}},
	}

	for _, test := range tests {
		result := MoveZeroEnd(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestMoveZeroEndUsingSwap(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0, 0, 1}, []int{1, 0, 0}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{1, 0, 2, 0, 3}, []int{1, 2, 3, 0, 0}},
	}

	for _, test := range tests {
		result := MoveZeroEndUsingSwap(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func TestRotateByKInLeft(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{3, 4, 5, 1, 2}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 0, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{2, 3, 4, 5, 1}},
		{[]int{1, 2, 3, 4, 5}, 7, []int{3, 4, 5, 1, 2}}, // k > len(input)
	}

	for _, test := range tests {
		result := RotateByKInLeft(test.input, test.k)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v and k %d, expected %v, but got %v", test.input, test.k, test.expected, result)
		}
	}
}

func TestRotateByTempArray(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{3, 4, 5, 1, 2}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 0, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{2, 3, 4, 5, 1}},
		{[]int{1, 2, 3, 4, 5}, 7, []int{3, 4, 5, 1, 2}}, // k > len(input)
	}

	for _, test := range tests {
		result := RotateByTempArray(test.input, test.k)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v and k %d, expected %v, but got %v", test.input, test.k, test.expected, result)
		}
	}
}


func TestRotateByReversal(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{3, 4, 5, 1, 2}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 0, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{2, 3, 4, 5, 1}},
		{[]int{1, 2, 3, 4, 5}, 7, []int{3, 4, 5, 1, 2}}, // k > len(input)
	}

	for _, test := range tests {
		result := RotateByReversal(test.input, test.k)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v and k %d, expected %v, but got %v", test.input, test.k, test.expected, result)
		}
	}
}

func TestRotateByKInRight(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 4, 5, 1, 2}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 0, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{5, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 7, []int{4, 5, 1, 2, 3}}, // k > len(input)
	}

	for _, test := range tests {
		result := RotateByKInRight(test.input, test.k)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v and k %d, expected %v, but got %v", test.input, test.k, test.expected, result)
		}
	}

}

func TestRotateByTempRightArray(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 4, 5, 1, 2}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 0, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{5, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 7, []int{4, 5, 1, 2, 3}}, // k > len(input)
	}

	for _, test := range tests {
		result := RotateByTempRightArray(test.input, test.k)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v and k %d, expected %v, but got %v", test.input, test.k, test.expected, result)
		}
	}

}
