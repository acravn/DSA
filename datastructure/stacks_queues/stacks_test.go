package stacks_queues

import (
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
