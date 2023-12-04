package stacks_queues

import "testing"

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
