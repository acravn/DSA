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
