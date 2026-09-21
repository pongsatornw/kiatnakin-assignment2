package main

import (
	"testing"

	"github.com/go-openapi/testify/assert"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		name     string
		slice1   []int
		slice2   []int
		slice3   []int
		expected []int
	}{
		{
			name:     "Merge 3 slices",
			slice1:   []int{15, 9, 5, 3},
			slice2:   []int{1, 2, 4},
			slice3:   []int{6, 8, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 15},
		},
		{
			name:     "Merge 3 slices, slice1 is empty",
			slice1:   []int{},
			slice2:   []int{1, 3, 5, 7},
			slice3:   []int{2, 4, 6, 8},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:     "Merge 3 slices, slice2 is empty",
			slice1:   []int{9, 7, 5, 3, 1},
			slice2:   []int{},
			slice3:   []int{2, 4, 6, 8},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "Merge 3 slices, slice3 is empty",
			slice1:   []int{9, 7, 5, 3},
			slice2:   []int{2, 4, 6, 8},
			slice3:   []int{},
			expected: []int{2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "Merge 3 slices, all slices are empty",
			slice1:   []int{},
			slice2:   []int{},
			slice3:   []int{},
			expected: []int{},
		},
		{
			name:     "Merge 3 slices, all slices are nil",
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Merge(tc.slice1, tc.slice2, tc.slice3)

			assert.EqualValues(t, tc.expected, actual)
		})
	}
}
