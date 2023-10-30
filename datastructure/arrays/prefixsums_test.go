package arrays

import (
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {
	table := []struct {
		arr      []int
		expected []int
	}{
		{
			arr:      []int{1, 2, 3, 4},
			expected: []int{1, 3, 6, 10},
		},
		{
			arr:      []int{1, 1, 1, 1, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			arr:      []int{3, 1, 2, 10, 1},
			expected: []int{3, 4, 6, 16, 17},
		},
	}

	for _, v := range table {
		res := runningSum(v.arr)
		if !reflect.DeepEqual(res, v.expected) {
			t.Errorf("Got %v, expected %v", res, v.expected)
		}
	}
}

func TestMinStartValue(t *testing.T) {
	table := []struct {
		arr      []int
		expected int
	}{
		{
			arr:      []int{-3, 2, -3, 4, 2},
			expected: 5,
		},
		{
			arr:      []int{1, 2},
			expected: 1,
		},
		{
			arr:      []int{1, -2, -3},
			expected: 5,
		},
	}

	for _, v := range table {
		res := minStartValue(v.arr)
		if res != v.expected {
			t.Errorf("Got %v, expected %v", res, v.expected)
		}
	}
}
