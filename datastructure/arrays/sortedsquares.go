package arrays

func SortedSquares(nums []int) []int {
	l, r, k := 0, len(nums)-1, len(nums)-1

	result := make([]int, len(nums))

	for l <= r {
		l2, r2 := nums[l]*nums[l], nums[r]*nums[r]
		if l2 > r2 {
			result[k] = l2
			l++
		} else {
			result[k] = r2
			r--
		}
		k--
	}

	return result
}
