package arrays

// Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.
func findDisappearedNumbers(nums []int) []int {
	mem := make([]bool, len(nums))
	for _, x := range nums {
		mem[x-1] = true
	}
	res := []int{}
	for i, x := range mem {
		if !x {
			res = append(res, i+1)
		}
	}
	return res
}
