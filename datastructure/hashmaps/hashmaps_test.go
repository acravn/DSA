package hashmaps

import "testing"

func TestPangram(t *testing.T) {
	table := []struct {
		input    string
		expected bool
	}{
		{
			input:    "thequickbrownfoxjumpsoverthelazydog",
			expected: true,
		},
		{
			input:    "leetcode",
			expected: false,
		},
	}

	for _, v := range table {
		res := checkIfPangram(v.input)
		if res != v.expected {
			t.Errorf("Got %t, expected %t", res, v.expected)
		}
	}
}

func TestCountElements(t *testing.T) {
	table := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 2, 3},
			expected: 2,
		},
		{
			input:    []int{1, 1, 3, 3, 5, 5, 7, 7},
			expected: 0,
		},
		{
			input:    []int{1, 3, 2, 3, 5, 0},
			expected: 3,
		},
	}

	for _, v := range table {
		res := countElements(v.input)
		if res != v.expected {
			t.Errorf("Expected %d, got %d", v.expected, res)
		}
	}
}
