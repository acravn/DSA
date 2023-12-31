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

// Given a string text, you want to use the characters of text to form as many instances of the word "balloon" as possible.
// You can use each character in text at most once. Return the maximum number of instances that can be formed.
func maxNumberOfBalloons(text string) int {
	letterMap := make(map[string]int, 5)
	for _, v := range text {
		letter := string(v)

		if letter == "b" {
			letterMap["b"] += 1
		}
		if letter == "a" {
			letterMap["a"] += 1
		}
		if letter == "l" {
			letterMap["l"] += 1
		}
		if letter == "o" {
			letterMap["o"] += 1
		}
		if letter == "n" {
			letterMap["n"] += 1
		}

	}

	count := letterMap["b"]
	if letterMap["a"] < count {
		count = letterMap["a"]
	}
	if letterMap["l"] < count*2 {
		count = letterMap["l"] / 2
	}
	if letterMap["o"] < count*2 {
		count = letterMap["o"] / 2
	}
	if letterMap["n"] < count {
		count = letterMap["n"]
	}
	return count
}

// Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
// Each letter in magazine can only be used once in ransomNote.
func canConstruct(ransomNote string, magazine string) bool {

	if len(ransomNote) > len(magazine) {
		return false
	}

	magazineMap := make(map[string]int, len(magazine))
	for _, v := range magazine {
		letter := string(v)
		if _, ok := magazineMap[letter]; !ok {
			magazineMap[letter] = 1
		} else {
			magazineMap[letter] += 1
		}
	}

	for _, v := range ransomNote {
		letter := string(v)

		if val, ok := magazineMap[letter]; !ok {
			return false
		} else if val > 0 {
			magazineMap[letter] -= 1
		} else {
			delete(magazineMap, letter)
			return false
		}

	}

	return true
}

// You're given strings jewels representing the types of stones that are jewels, and stones representing the stones you have.
// Each character in stones is a type of stone you have. You want to know how many of the stones you have are also jewels.
// Letters are case sensitive, so "a" is considered a different type of stone from "A".
func numJewelsInStones(jewels string, stones string) int {
	stoneCounter := make(map[string]int, len(stones))
	var counter int
	for _, v := range stones {
		letter := string(v)
		if _, ok := stoneCounter[letter]; !ok {
			stoneCounter[letter] = 1
		} else {
			stoneCounter[letter] += 1
		}
	}

	for _, v := range jewels {
		letter := string(v)
		val, ok := stoneCounter[letter]
		if ok {
			counter += val
		}
	}

	return counter
}
