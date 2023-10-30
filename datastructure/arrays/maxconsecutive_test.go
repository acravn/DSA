package arrays

import (
	"testing"
)

func TestMaxConsecutive(t *testing.T) {
	table := []struct {
		arr      []int
		k        int
		expected int
	}{
		//{
		//	arr:      []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
		//	k:        2,
		//	expected: 6,
		//},
		{
			arr:      []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			k:        3,
			expected: 10,
		},
	}

	for _, v := range table {
		res := longestOnes(v.arr, v.k)
		if res != v.expected {
			t.Errorf("Got %v, expected %v", res, v.expected)
		}
	}
}
