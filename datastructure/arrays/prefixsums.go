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
