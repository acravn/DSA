package lc

import "log"

// Given the head of a singly linked list, return the middle node of the linked list.
// If there are two middle nodes, return the second middle node.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// Given the head of a sorted linked list, delete all duplicates such that each
// element appears only once. Return the linked list sorted as well.
func deleteDuplicates(head *ListNode) *ListNode {
	res := head
	for head != nil {
		for head.Next != nil && head.Next.Val == head.Val {
			head.Next = head.Next.Next
		}
		head = head.Next
	}
	return res
}

func printLL(head *ListNode) {
	for {
		log.Println("Val: ", head.Val, "Next: ", head.Next)
		if head.Next == nil {
			break
		}
		head = head.Next
	}
}

// reverse a linked list
func reverseLL(head *ListNode) *ListNode {
	node := head
	var prev *ListNode
	for node != nil {
		node, prev, node.Next = node.Next, node, prev
	}

	head = prev

	return head
}

// Given the head of a singly linked list and two integers left and right where left <= right,
// reverse the nodes of the list from position left to position right, and return the reversed list.
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{0, head}
	prev := dummy

	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	current := prev.Next

	for i := 0; i < right-left; i++ {
		nextNode := current.Next
		current.Next = nextNode.Next
		nextNode.Next = prev.Next
		prev.Next = nextNode
	}

	return dummy.Next
}
