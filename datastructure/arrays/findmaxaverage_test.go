package arrays

import (
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	table := []struct {
		arr      []int
		k        int
		expected float64
	}{
		{
			arr:      []int{1, 12, -5, -6, 50, 3},
			k:        4,
			expected: 12.75000,
		},
	}

	for _, v := range table {
		res := FindMaxAverage(v.arr, v.k)
		if res != v.expected {
			t.Errorf("Got %v, expected %v", res, v.expected)
		}
	}
}
