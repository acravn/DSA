package twocrystal

import "testing"

func TestTwoCrystalBalls(t *testing.T) {
	table := []struct {
		building []bool
		expected int
	}{
		{
			building: []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true},
			expected: 14,
		},
	}

	for _, v := range table {
		res := TwoCrystalBalls(v.building)

		if res != v.expected {
			t.Errorf("Got %d, Expected: %v", res, v.expected)
		}
	}
}

func BenchmarkTwoCrystalBalls(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoCrystalBalls([]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true})
	}
}
