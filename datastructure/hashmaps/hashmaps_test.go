package hashmaps

import (
	"reflect"
	"testing"
)

func TestPangram(t *testing.T) {
	table := []struct {
		input    string
		expected bool
	}{
		{
			input:    "thequickbrownfoxjumpsoverthelazydog",
			expected: true,
		},
		{
			input:    "leetcode",
			expected: false,
		},
	}

	for _, v := range table {
		res := checkIfPangram(v.input)
		if res != v.expected {
			t.Errorf("Got %t, expected %t", res, v.expected)
		}
	}
}

func TestCountElements(t *testing.T) {
	table := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 2, 3},
			expected: 2,
		},
		{
			input:    []int{1, 1, 3, 3, 5, 5, 7, 7},
			expected: 0,
		},
		{
			input:    []int{1, 3, 2, 3, 5, 0},
			expected: 3,
		},
	}

	for _, v := range table {
		res := countElements(v.input)
		if res != v.expected {
			t.Errorf("Expected %d, got %d", v.expected, res)
		}
	}
}

func TestFindWinners(t *testing.T) {
	tables := []struct {
		input  [][]int
		output [][]int
	}{
		{
			input:  [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}},
			output: [][]int{{1, 2, 10}, {4, 5, 7, 8}},
		},
	}

	for _, test := range tables {
		res := findWinners(test.input)

		if !reflect.DeepEqual(res, test.output) {
			t.Errorf("Expected %v, got %v", test.output, res)
		}
	}
}

func TestLargestUniqueNumber(t *testing.T) {
	tables := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{5, 7, 3, 9, 4, 9, 8, 3, 1},
			output: 8,
		},
		{
			input:  []int{9, 9, 8, 8},
			output: -1,
		},
		{
			input:  []int{9, 8, 8, 9, 9},
			output: -1,
		},
	}

	for _, test := range tables {
		res := largestUniqueNumber(test.input)

		if res != test.output {
			t.Errorf("Expected %v, got %v", test.output, res)
		}
	}
}

func TestMaxNumBalloons(t *testing.T) {
	tables := []struct {
		input  string
		output int
	}{
		{
			input:  "nlaebolko",
			output: 1,
		},
		//{
		//	input:  "loonbalxballpoon",
		//	output: 2,
		//},
		//{
		//	input:  "leetcode",
		//	output: 0,
		//},
	}

	for _, test := range tables {
		res := maxNumberOfBalloons(test.input)

		if res != test.output {
			t.Errorf("Expected %v, got %v", test.output, res)
		}
	}
}

func TestCanConstruct(t *testing.T) {
	tables := []struct {
		inputA string
		inputB string
		output bool
	}{
		{
			inputA: "a",
			inputB: "b",
			output: false,
		},
		{
			inputA: "aa",
			inputB: "ab",
			output: false,
		},
		{
			inputA: "aa",
			inputB: "aab",
			output: true,
		},
	}

	for _, test := range tables {
		res := canConstruct(test.inputA, test.inputB)

		if res != test.output {
			t.Errorf("Expected %v, got %v", test.output, res)
		}
	}
}

func TestNumJewelsInStones(t *testing.T) {
	tables := []struct {
		inputA string
		inputB string
		output int
	}{
		{
			inputA: "aA",
			inputB: "aAAbbbb",
			output: 3,
		},
		{
			inputA: "z",
			inputB: "ZZ",
			output: 0,
		},
	}

	for _, test := range tables {
		res := numJewelsInStones(test.inputA, test.inputB)

		if res != test.output {
			t.Errorf("Expected %v, got %v", test.output, res)
		}
	}
}
