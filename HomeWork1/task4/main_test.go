package task4

import (
	"reflect"
	"testing"
)

func TestSortBinaryArray(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected []int
	}{
		{
			name:     "mixed values",
			arr:      []int{1, 0, 1, 0, 0, 1},
			expected: []int{0, 0, 0, 1, 1, 1},
		},
		{
			name:     "empty array",
			arr:      []int{},
			expected: []int{},
		},
		{
			name:     "single zero",
			arr:      []int{0},
			expected: []int{0},
		},
		{
			name:     "single one",
			arr:      []int{1},
			expected: []int{1},
		},
		{
			name:     "all zeroes",
			arr:      []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "all ones",
			arr:      []int{1, 1, 1},
			expected: []int{1, 1, 1},
		},
		{
			name:     "already sorted",
			arr:      []int{0, 0, 1, 1},
			expected: []int{0, 0, 1, 1},
		},
		{
			name:     "reverse sorted",
			arr:      []int{1, 1, 0, 0},
			expected: []int{0, 0, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortBinaryArray(tt.arr)

			if !reflect.DeepEqual(tt.arr, tt.expected) {
				t.Errorf("sortBinaryArray() = %v, want %v", tt.arr, tt.expected)
			}
		})
	}
}
