package task1

import (
	"reflect"
	"testing"
)

func TestReversePartOfArray(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		k        int
		expected []int
	}{
		{
			name:     "rotate by 3",
			arr:      []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:     "rotate by 0",
			arr:      []int{1, 2, 3, 4, 5},
			k:        0,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "rotate by array length",
			arr:      []int{1, 2, 3},
			k:        3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "rotate more than array length",
			arr:      []int{1, 2, 3},
			k:        4,
			expected: []int{3, 1, 2},
		},
		{
			name:     "empty array",
			arr:      []int{},
			k:        3,
			expected: []int{},
		},
		{
			name:     "single element",
			arr:      []int{1},
			k:        5,
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reversePartOfArray(tt.arr, tt.k)

			if !reflect.DeepEqual(tt.arr, tt.expected) {
				t.Errorf("reversePartOfArray() = %v, want %v",
					tt.arr, tt.expected)
			}
		})
	}
}
