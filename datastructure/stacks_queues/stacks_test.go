package stacks_queues

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		//{
		//	input:  "/home/",
		//	output: "/home",
		//},
		{
			input:  "/../",
			output: "/",
		},
		//{
		//	input:  "/home//foo/",
		//	output: "/home/foo",
		//},
		{
			input:  "/a/./b/../../c/",
			output: "/c",
		},
	}

	for _, v := range tests {
		res := simplifyPath(v.input)
		if res != v.output {
			t.Errorf("Expected %s, got %s", v.output, res)
		}
	}
}

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		inputA []int
		inputB []int
		output []int
	}{
		{
			inputA: []int{4, 1, 2},
			inputB: []int{1, 3, 4, 2},
			output: []int{-1, 3, -1},
		},
		{
			inputA: []int{2, 4},
			inputB: []int{1, 2, 3, 4},
			output: []int{3, -1},
		},
	}

	for _, v := range tests {
		res := nextGreaterElement(v.inputA, v.inputB)
		if !reflect.DeepEqual(res, v.output) {
			t.Errorf("Expected %v, got %v", v.output, res)
		}
	}
}

func TestStockSpanner_Next(t *testing.T) {
	cases := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{100, 80, 60, 70, 60, 75, 85},
			output: []int{1, 1, 1, 2, 1, 4, 6},
		},
	}

	obj := Constructor()

	for _, test := range cases {
		for i, price := range test.input {
			res := obj.Next(price)
			fmt.Printf("Response for price %d is %d\n", price, res)
			if res != test.output[i] {
				t.Errorf("Expected %v, got %v", test.output[i], res)
			}
		}
	}
}
