package binarysearch

import (
	"fmt"
	"math"
)

func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)
	for low < high {
		f32Low := float64(low)
		f32hi := float64(high)
		midpoint := int(math.Floor(f32Low + (f32hi-f32Low)/2))
		value := arr[midpoint]
		fmt.Println("midpoint:", midpoint, "- value:", value, "- target:", target, "- low:", low, "- high:", high)
		if value == target {
			return midpoint
		} else if value > target {
			high = midpoint
		} else {
			low = midpoint + 1
		}
	}
	return -1
}
