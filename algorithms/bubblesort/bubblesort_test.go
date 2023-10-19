package bubblesort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	table := []struct {
		arr      []int
		expected []int
	}{
		{
			arr:      []int{1, 4, 2, 6, 9, 7, 8},
			expected: []int{1, 2, 4, 6, 7, 8, 9},
		},
	}

	for _, v := range table {
		BubbleSort(v.arr)
		if !reflect.DeepEqual(v.arr, v.expected) {
			t.Errorf("Got %v, expected %v", v.arr, v.expected)
		}
	}
}
