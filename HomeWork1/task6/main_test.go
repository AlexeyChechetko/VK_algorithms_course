package task6

import (
	"reflect"
	"testing"
)

func TestEvenFirst(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected []int
	}{
		{
			name:     "move evens forward",
			arr:      []int{3, 2, 4, 1, 11, 8, 9},
			expected: []int{2, 4, 8, 1, 11, 3, 9},
		},
		{
			name:     "empty array",
			arr:      []int{},
			expected: []int{},
		},
		{
			name:     "single even",
			arr:      []int{2},
			expected: []int{2},
		},
		{
			name:     "single odd",
			arr:      []int{3},
			expected: []int{3},
		},
		{
			name:     "all even",
			arr:      []int{8, 2, 4},
			expected: []int{8, 2, 4},
		},
		{
			name:     "all odd",
			arr:      []int{3, 1, 5},
			expected: []int{3, 1, 5},
		},
		{
			name:     "already partitioned",
			arr:      []int{2, 4, 3, 1},
			expected: []int{2, 4, 3, 1},
		},
		{
			name:     "negative numbers and zero",
			arr:      []int{-3, -2, 0, -4, 5},
			expected: []int{-2, 0, -4, -3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evenFirst(tt.arr)

			if !reflect.DeepEqual(tt.arr, tt.expected) {
				t.Errorf("evenFirst() = %v, want %v", tt.arr, tt.expected)
			}
		})
	}
}
