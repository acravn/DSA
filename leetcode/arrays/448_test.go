package arrays

import (
	"reflect"
	"testing"
)

func Test_findDisappearedNumbers(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{4, 3, 2, 7, 8, 2, 3, 1},
			expected: []int{5, 6},
		},
		//{
		//	input:    []int{1, 1},
		//	expected: []int{2},
		//},
	}

	for _, v := range tests {
		res := findDisappearedNumbers(v.input)
		if !reflect.DeepEqual(res, v.expected) {
			t.Errorf("Got: %v, Expected: %v", res, v.expected)
		}
	}
}
