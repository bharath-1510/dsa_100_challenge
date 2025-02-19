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
