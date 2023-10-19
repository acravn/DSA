package main

import (
	"fmt"
	"math"
)

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

func BubbleSort(arr []int) {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

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

func main() {
	building := []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true}

	fmt.Println(TwoCrystalBalls(building))

	bubbleArr := []int{1, 4, 2, 6, 9, 7, 8}
	fmt.Println(bubbleArr)
	BubbleSort(bubbleArr)
	fmt.Println(BinarySearch(bubbleArr, 9))
	fmt.Println(bubbleArr)

}
