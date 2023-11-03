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
