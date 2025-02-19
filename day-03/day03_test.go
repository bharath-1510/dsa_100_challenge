package day03

import "testing"

func TestIntersectTwoArray(t *testing.T) {
	tests := []struct {
		name     string
		array1   []int
		array2   []int
		expected []int
	}{
		{
			name:     "No intersection",
			array1:   []int{1, 2, 3},
			array2:   []int{4, 5, 6},
			expected: []int{},
		},
		{
			name:     "Some intersection",
			array1:   []int{1, 2, 3, 4},
			array2:   []int{3, 4, 5, 6},
			expected: []int{3, 4},
		},
		{
			name:     "Full intersection",
			array1:   []int{1, 2, 3},
			array2:   []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Empty arrays",
			array1:   []int{},
			array2:   []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IntersectTwoArray(tt.array1, tt.array2)
			if !equal(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestIntersectTwoArrayUsingMap(t *testing.T) {
	tests := []struct {
		name     string
		array1   []int
		array2   []int
		expected []int
	}{
		{
			name:     "No intersection",
			array1:   []int{1, 2, 3},
			array2:   []int{4, 5, 6},
			expected: []int{},
		},
		{
			name:     "Some intersection",
			array1:   []int{1, 2, 3, 4},
			array2:   []int{3, 4, 5, 6},
			expected: []int{3, 4},
		},
		{
			name:     "Full intersection",
			array1:   []int{1, 2, 3},
			array2:   []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Empty arrays",
			array1:   []int{},
			array2:   []int{},
			expected: []int{},
		},
		{
			name:     "One empty array",
			array1:   []int{1, 2, 3},
			array2:   []int{},
			expected: []int{},
		},
		{
			name:     "Identical arrays with duplicates",
			array1:   []int{1, 2, 2, 3, 3},
			array2:   []int{1, 2, 2, 3, 3},
			expected: []int{1, 2, 2, 3, 3},
		},
		{
			name:     "Arrays with different lengths",
			array1:   []int{1, 2, 3, 4, 5},
			array2:   []int{3, 4},
			expected: []int{3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IntersectTwoArrayWithDuplicates(tt.array1, tt.array2)
			if !equal(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIntersectTwoArrayWithDuplicatesMerge(t *testing.T) {
	tests := []struct {
		name     string
		array1   []int
		array2   []int
		expected []int
	}{
		{
			name:     "No intersection",
			array1:   []int{1, 2, 3},
			array2:   []int{4, 5, 6},
			expected: []int{},
		},
		{
			name:     "Some intersection",
			array1:   []int{1, 2, 3, 4},
			array2:   []int{3, 4, 5, 6},
			expected: []int{3, 4},
		},
		{
			name:     "Full intersection",
			array1:   []int{1, 2, 3},
			array2:   []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Empty arrays",
			array1:   []int{},
			array2:   []int{},
			expected: []int{},
		},
		{
			name:     "One empty array",
			array1:   []int{1, 2, 3},
			array2:   []int{},
			expected: []int{},
		},
		{
			name:     "Identical arrays with duplicates",
			array1:   []int{1, 2, 2, 3, 3},
			array2:   []int{1, 2, 2, 3, 3},
			expected: []int{1, 2, 2, 3, 3},
		},
		{
			name:     "Arrays with different lengths",
			array1:   []int{1, 2, 3, 4, 5},
			array2:   []int{3, 4},
			expected: []int{3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IntersectTwoArrayWithDuplicatesMerge(tt.array1, tt.array2)
			if !equal(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIntersectTwoArrayWithoutDuplicatesMerge(t *testing.T) {
	tests := []struct {
		name     string
		array1   []int
		array2   []int
		expected []int
	}{
		{
			name:     "No intersection",
			array1:   []int{1, 2, 3},
			array2:   []int{4, 5, 6},
			expected: []int{},
		},
		{
			name:     "Some intersection",
			array1:   []int{1, 2, 3, 4},
			array2:   []int{3, 4, 5, 6},
			expected: []int{3, 4},
		},
		{
			name:     "Full intersection",
			array1:   []int{1, 2, 3},
			array2:   []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Empty arrays",
			array1:   []int{},
			array2:   []int{},
			expected: []int{},
		},
		{
			name:     "One empty array",
			array1:   []int{1, 2, 3},
			array2:   []int{},
			expected: []int{},
		},
		{
			name:     "Identical arrays with duplicates",
			array1:   []int{1, 2, 2, 3, 3},
			array2:   []int{1, 2, 2, 3, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Arrays with different lengths",
			array1:   []int{1, 2, 3, 4, 5},
			array2:   []int{3, 4},
			expected: []int{3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IntersectTwoArrayWithoutDuplicatesMerge(tt.array1, tt.array2)
			if !equal(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIsArraySorted(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		expected bool
	}{
		{
			name:     "Sorted array",
			array:    []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "Unsorted array",
			array:    []int{5, 3, 1, 2, 4},
			expected: false,
		},
		{
			name:     "Empty array",
			array:    []int{},
			expected: true,
		},
		{
			name:     "Single element array",
			array:    []int{1},
			expected: true,
		},
		{
			name:     "Array with duplicates",
			array:    []int{1, 2, 2, 3, 4},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsArraySorted(tt.array)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestFindMissingNumber(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		expected int
	}{
		
		{
			name:     "Missing number in the middle",
			array:    []int{1, 2, 4, 5},
			expected: 3,
		},
		{
			name:     "Missing number at the start",
			array:    []int{2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "Missing number at the end",
			array:    []int{1, 2, 3, 4},
			expected: 5,
		},
		{
			name:     "Empty array",
			array:    []int{},
			expected: 1,
		},
		{
			name:     "Single element array",
			array:    []int{2},
			expected: 1,
		},
		{
			name:     "Array with duplicates",
			array:    []int{1, 2, 2, 4, 5},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindMissingNumber(tt.array)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestFindMissingNumberUsingFormula(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		expected int
	}{
		
		{
			name:     "Missing number in the middle",
			array:    []int{1, 2, 4, 5},
			expected: 3,
		},
		{
			name:     "Missing number at the start",
			array:    []int{2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "Missing number at the end",
			array:    []int{1, 2, 3, 4},
			expected: 5,
		},
		{
			name:     "Empty array",
			array:    []int{},
			expected: 1,
		},
		{
			name:     "Single element array",
			array:    []int{2},
			expected: 1,
		},
		{
			name:     "Array with duplicates",
			array:    []int{1, 2, 2, 4, 5},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindMissingNumberUsingFormula(tt.array)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestFindMissingNumberUsingHash(t *testing.T) {
	tests := []struct {
		name     string
		array    []int
		expected int
	}{
		
		{
			name:     "Missing number in the middle",
			array:    []int{1, 2, 4, 5},
			expected: 3,
		},
		{
			name:     "Missing number at the start",
			array:    []int{2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "Missing number at the end",
			array:    []int{1, 2, 3, 4},
			expected: 5,
		},
		{
			name:     "Empty array",
			array:    []int{},
			expected: 1,
		},
		{
			name:     "Single element array",
			array:    []int{2},
			expected: 1,
		},
		{
			name:     "Array with duplicates",
			array:    []int{1, 2, 2, 4, 5},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindMissingNumberUsingHash(tt.array)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}

}
