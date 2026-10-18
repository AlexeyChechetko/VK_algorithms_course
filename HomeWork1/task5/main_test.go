package task5

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected []int
	}{
		{
			name:     "mixed colors",
			arr:      []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
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
			name:     "single two",
			arr:      []int{2},
			expected: []int{2},
		},
		{
			name:     "all same",
			arr:      []int{2, 2, 2},
			expected: []int{2, 2, 2},
		},
		{
			name:     "already sorted",
			arr:      []int{0, 0, 1, 1, 2, 2},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "reverse sorted",
			arr:      []int{2, 2, 1, 1, 0, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "recheck swapped value",
			arr:      []int{1, 2, 0},
			expected: []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.arr)

			if !reflect.DeepEqual(tt.arr, tt.expected) {
				t.Errorf("sortColors() = %v, want %v", tt.arr, tt.expected)
			}
		})
	}
}
