package arrays

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	table := []struct {
		arr      []int
		expected []int
	}{
		{
			arr:      []int{-4, -1, 0, 3, 10},
			expected: []int{0, 1, 9, 16, 100},
		},
		{
			arr:      []int{-7, -3, 2, 3, 11},
			expected: []int{4, 9, 9, 49, 121},
		},
	}

	for _, v := range table {
		res := SortedSquares(v.arr)
		if !reflect.DeepEqual(res, v.expected) {
			t.Errorf("Got %v, expected %v", res, v.expected)
		}
	}
}
