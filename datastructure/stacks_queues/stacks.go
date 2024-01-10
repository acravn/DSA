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

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	hashMap := make(map[int]int, len(nums1))
	var stack []int
	for _, v := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < v {
			hashMap[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}
	for i, v := range nums1 {
		if mapv, ok := hashMap[v]; ok {
			nums1[i] = mapv
			continue
		}
		nums1[i] = -1
	}
	return nums1
}
