package arrays

// Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
// Return the running sum of nums.
func runningSum(nums []int) []int {
	prefix := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			prefix[i] = prefix[i-1] + nums[i]
		} else {
			prefix[i] = nums[i]
		}
	}
	return prefix
}

// Given an array of integers nums, you start with an initial positive value startValue.
// In each iteration, you calculate the step by step sum of startValue plus elements in nums (from left to right).
// Return the minimum positive value of startValue such that the step by step sum is never less than 1.
func minStartValue(nums []int) int {
	var minVal, sum int
	for _, v := range nums {
		sum += v
		if sum < minVal {
			minVal = sum
		}
	}

	if minVal < 0 {
		minVal = -minVal
	}

	return minVal + 1
}

// You are given a 0-indexed array nums of n integers, and an integer k.
// The k-radius average for a subarray of nums centered at some index i with the radius k
// is the average of all elements in nums between the indices i - k and i + k (inclusive).
// If there are less than k elements before or after the index i, then the k-radius average is -1.
// Build and return an array avgs of length n where avgs[i] is the k-radius average for the subarray centered at index i.
// The average of x elements is the sum of the x elements divided by x, using integer division.
// The integer division truncates toward zero, which means losing its fractional part.
func getAverages(nums []int, k int) []int {
	numsLen := len(nums)
	avgs := make([]int, numsLen)
	prefixSum := make([]int, numsLen)
	windowSize := k*2 + 1

	prefixSum[0] = nums[0]
	for i := 1; i < numsLen; i++ {
		prefixSum[i] = nums[i] + prefixSum[i-1]
	}

	for i, _ := range prefixSum {
		left, right := i-k, i+k
		if left < 0 || right > numsLen-1 {
			avgs[i] = -1
		} else if k == 0 {
			avgs[i] = nums[i]
		} else if left > 0 {
			avgs[i] = (prefixSum[right] - prefixSum[left-1]) / windowSize
		} else {
			avgs[i] = prefixSum[right] / windowSize
		}

	}

	return avgs
}
