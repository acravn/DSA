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

func TestGetAverages(t *testing.T) {
	table := []struct {
		arr      []int
		k        int
		expected []int
	}{
		{
			arr:      []int{7, 4, 3, 9, 1, 8, 5, 2, 6}, // {7, 11, 14, 23, 24, 32, 37, 39, 45}
			k:        3,
			expected: []int{-1, -1, -1, 5, 4, 4, -1, -1, -1},
		},
		{
			arr:      []int{100000},
			k:        0,
			expected: []int{100000},
		},
		{
			arr:      []int{8},
			k:        10000,
			expected: []int{-1},
		},
		{
			arr:      []int{18334, 25764, 19780, 92480, 69842, 73255, 89893},
			k:        0,
			expected: []int{18334, 25764, 19780, 92480, 69842, 73255, 89893},
		},
		{
			arr:      []int{40527, 53696, 10730, 66491, 62141, 83909, 78635, 18560},
			k:        2,
			expected: []int{-1, -1, 46717, 55393, 60381, 61947, -1, -1},
		},
	}

	for _, v := range table {
		res := getAverages(v.arr, v.k)
		if !reflect.DeepEqual(res, v.expected) {
			t.Errorf("Got %v, expected %v", res, v.expected)
		}
	}
}
