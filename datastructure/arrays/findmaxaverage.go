package arrays

import (
	"math"
)

// Fixed sliding window problem
// You are given an integer array nums consisting of n elements, and an integer k.
// Find a contiguous subarray whose length is equal to k that has the maximum average
// value and return this value. Any answer with a calculation error less than 10-5 will be accepted.
func FindMaxAverage(nums []int, k int) float64 {
	var currSum int
	var avg float64
	// build initial window of size k, find it's average
	for i := 0; i < k; i++ {
		currSum += nums[i]
	}
	avg = float64(currSum) / float64(k)
	// start sliding the window, compare to previous average and keep largest of two
	for i := k; i < len(nums); i++ {
		currSum += nums[i]
		currSum -= nums[i-k]
		avg = math.Max(avg, float64(currSum)/float64(k))
	}
	return avg
}
