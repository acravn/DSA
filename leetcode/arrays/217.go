package arrays

func containsDuplicate(nums []int) bool {
	keys := map[int]bool{}
	for _, v := range nums {
		if keys[v] {
			return true
		}
		keys[v] = true
	}
	return false
}
