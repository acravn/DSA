package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	table := []struct {
		haystack []int
		needle   int
		expected int
	}{
		{
			haystack: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
			needle:   2,
			expected: 1,
		},
		{
			haystack: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
			needle:   9,
			expected: 8,
		},
	}

	for _, v := range table {
		res := BinarySearch(v.haystack, v.needle)

		if res != v.expected {
			t.Errorf("Got: %d, Expected: %d\n", res, v.expected)
		}
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}, 2)
	}
}
