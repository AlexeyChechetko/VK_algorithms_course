package task7

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected []int
	}{
		{
			name:     "move zeroes to end",
			arr:      []int{0, 0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0, 0},
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
			name:     "single nonzero",
			arr:      []int{3},
			expected: []int{3},
		},
		{
			name:     "all zeroes",
			arr:      []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "no zeroes",
			arr:      []int{3, 1, 2},
			expected: []int{3, 1, 2},
		},
		{
			name:     "zeroes already at end",
			arr:      []int{1, 2, 0, 0},
			expected: []int{1, 2, 0, 0},
		},
		{
			name:     "negative values",
			arr:      []int{0, -3, 0, -1, 2},
			expected: []int{-3, -1, 2, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.arr)

			if !reflect.DeepEqual(tt.arr, tt.expected) {
				t.Errorf("moveZeroes() = %v, want %v", tt.arr, tt.expected)
			}
		})
	}
}
