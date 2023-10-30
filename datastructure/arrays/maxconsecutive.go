package arrays

// dynamic sliding window
// Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.
func longestOnes(nums []int, k int) int {
	var left, right, flipCount, windowSize, ans int

	for i := 0; i < len(nums); i++ {
		// increment window
		right = i
		if nums[i] == 0 {
			flipCount++
		}

		// if number of 0s exceeds k limit, start shifting the window from the left until enough 0s have been removed to start growing window again
		for flipCount > k {
			if nums[left] == 0 {
				flipCount--
			}
			left++
		}
		windowSize = right - left + 1

		// check for max without using math.Max that requires float64 type
		if windowSize > ans {
			ans = windowSize
		}
	}
	return ans
}
