package stacks_queues

import (
	"strings"
)

func simplifyPath(path string) string {
	stack := make([]string, 0)
	split := strings.Split(path, "/")
	for _, v := range split {
		if v == "" || v == "." {
			continue
		}
		if v != ".." {
			stack = append(stack, v)
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}

	return "/" + strings.Join(stack, "/")
}
