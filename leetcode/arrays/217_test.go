package arrays

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{
			input:    []int{1, 2, 3, 1},
			expected: true,
		},
		{
			input:    []int{1, 2, 3, 4},
			expected: false,
		},
		{
			input:    []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
	}

	for _, v := range tests {
		res := containsDuplicate(v.input)
		if res != v.expected {
			t.Errorf("Got: %v, Expected: %v", res, v.expected)
		}
	}
}
