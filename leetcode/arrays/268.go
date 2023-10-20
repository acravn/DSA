package arrays

// Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.
func missingNumber(nums []int) int {
	n := len(nums)
	expectedTotal := (n * (n + 1)) / 2
	var actualTotal int
	for _, v := range nums {
		actualTotal += v
	}
	return expectedTotal - actualTotal
}
