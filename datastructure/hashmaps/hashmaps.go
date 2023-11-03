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
