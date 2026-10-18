package task3

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		arr1     []int
		arr2     []int
		expected []int
	}{
		{
			name:     "merge into first array",
			arr1:     []int{1, 2, 3, 0, 0, 0},
			arr2:     []int{2, 5, 6},
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:     "empty arrays",
			arr1:     []int{},
			arr2:     []int{},
			expected: []int{},
		},
		{
			name:     "empty first data",
			arr1:     []int{0, 0},
			arr2:     []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "empty second array",
			arr1:     []int{1, 2},
			arr2:     []int{},
			expected: []int{1, 2},
		},
		{
			name:     "duplicates",
			arr1:     []int{1, 2, 2, 0, 0},
			arr2:     []int{2, 2},
			expected: []int{1, 2, 2, 2, 2},
		},
		{
			name:     "zero as data",
			arr1:     []int{0, 2, 0, 0},
			arr2:     []int{0, 3},
			expected: []int{0, 0, 2, 3},
		},
		{
			name:     "negative numbers",
			arr1:     []int{-5, -1, 0, 0},
			arr2:     []int{-3, 2},
			expected: []int{-5, -3, -1, 2},
		},
		{
			name:     "single elements",
			arr1:     []int{2, 0},
			arr2:     []int{1},
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.arr1, tt.arr2)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("merge() = %v, want %v", got, tt.expected)
			}
		})
	}
}
