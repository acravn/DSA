package arrays

import "testing"

func Test_missingNumber(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{3, 0, 1},
			expected: 2,
		},
		{
			input:    []int{0, 1},
			expected: 2,
		},
		{
			input:    []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			expected: 8,
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: 0,
		},
	}

	for _, v := range tests {
		res := missingNumber(v.input)
		if res != v.expected {
			t.Errorf("Got: %v, Expected: %v", res, v.expected)
		}
	}
}
