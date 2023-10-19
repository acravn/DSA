package twocrystal

import "math"

func TwoCrystalBalls(building []bool) int {
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(building)))))
	i := jumpAmount
	for {
		if i >= len(building) {
			break
		}
		if building[i] {
			break
		}
		i += jumpAmount
	}
	i -= jumpAmount

	for j := 0; j < jumpAmount && i < len(building); j += 1 {
		i += 1
		if building[i] {
			return i
		}

	}
	return -1
}
