package hashmaps

// A pangram is a sentence where every letter of the English alphabet appears at least once.
// Given a string sentence containing only lowercase English letters, return true if sentence is a pangram, or false otherwise.
func checkIfPangram(sentence string) bool {
	set := map[rune]bool{}

	for _, v := range sentence {
		_, ok := set[v]
		if !ok {
			set[v] = true
		}
	}

	return len(set) == 26
}

// Given an integer array arr, count how many elements x there are, such that x + 1 is also in arr.
// If there are duplicates in arr, count them separately.
func countElements(arr []int) int {
	var count int
	set := make(map[int]struct{})

	// make set first to account for unordered arr
	for _, v := range arr {
		_, ok := set[v]
		if !ok {
			set[v] = struct{}{}
		}
	}

	// iterate through arr and check for v+1 condition, if true, increment count
	for _, v := range arr {
		_, ok := set[v+1]
		if ok {
			count++
		}
	}

	return count
}
