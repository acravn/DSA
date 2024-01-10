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

type StockSpanner struct {
	stack [][]int
}

func Constructor() StockSpanner {
	return StockSpanner{stack: make([][]int, 0)}
}

func (this *StockSpanner) Next(price int) int {
	count := 1

	for len(this.stack) > 0 && this.stack[len(this.stack)-1][0] <= price {
		count += this.stack[len(this.stack)-1][1]
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, []int{price, count})
	return count
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
