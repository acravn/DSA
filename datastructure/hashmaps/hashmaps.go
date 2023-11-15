package hashmaps

import (
	"fmt"
	"sort"
)

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

// You are given an integer array matches where matches[i] = [winneri, loseri] indicates that
// the player winneri defeated player loseri in a match
// Return a list answer of size 2 where:
// answer[0] is a list of all players that have not lost any matches.
// answer[1] is a list of all players that have lost exactly one match.
func findWinners(matches [][]int) [][]int {
	res := make([][]int, 2)
	winnerMap := make(map[int]int, len(matches))
	loserMap := make(map[int]int, len(matches))

	for _, v := range matches {
		w := v[0]
		l := v[1]

		if _, ok := loserMap[l]; !ok {
			loserMap[l] = 1
		} else {
			loserMap[l] += 1
		}

		if _, ok := winnerMap[l]; ok {
			delete(winnerMap, l)
		}

		fmt.Println("Checking if", w, "in loserMap", loserMap)
		if _, ok := loserMap[w]; !ok {
			winnerMap[w] = 1
		} else {
			fmt.Println(w, "is in loserMap, deleting from winnerMap")
			delete(winnerMap, w)
		}

	}

	var winners []int
	var losers []int

	for k, _ := range winnerMap {
		winners = append(winners, k)
	}

	for k, v := range loserMap {
		if v == 1 {
			losers = append(losers, k)
		}
	}
	sort.Ints(winners)
	sort.Ints(losers)
	res[0] = winners
	res[1] = losers
	return res

}

// Given an integer array nums, return the largest integer that only occurs once. If no integer occurs once, return -1.
func largestUniqueNumber(nums []int) int {
	numCount := make(map[int]int, len(nums))

	for _, v := range nums {
		if _, ok := numCount[v]; ok {
			numCount[v] += 1
		} else {
			numCount[v] = 1
		}
	}

	l := -1
	for k, v := range numCount {
		if v == 1 && k > l {
			l = k
		}
	}

	return l
}
