package task2

import (
	"reflect"
	"testing"
)

func TestMergeSortedArrays(t *testing.T) {
	tests := []struct {
		name     string
		arr1     []int
		arr2     []int
		expected []int
	}{
		{
			name:     "merge sorted arrays",
			arr1:     []int{1, 3, 5},
			arr2:     []int{2, 4, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "empty arrays",
			arr1:     []int{},
			arr2:     []int{},
			expected: []int{},
		},
		{
			name:     "empty first array",
			arr1:     []int{},
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
			arr1:     []int{1, 2, 2},
			arr2:     []int{2, 3},
			expected: []int{1, 2, 2, 2, 3},
		},
		{
			name:     "negative numbers",
			arr1:     []int{-5, -1, 0},
			arr2:     []int{-3, 0, 2},
			expected: []int{-5, -3, -1, 0, 0, 2},
		},
		{
			name:     "single elements",
			arr1:     []int{2},
			arr2:     []int{1},
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeSortedArrays(tt.arr1, tt.arr2)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("mergeSortedArrays() = %v, want %v", got, tt.expected)
			}
		})
	}
}
