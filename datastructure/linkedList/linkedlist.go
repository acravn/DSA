package linkedList

import "fmt"

// Example interface
//type LinkedList interface {
//	Length()
//	InsertAt(item any, index int)
//	Remove(item any)
//	RemoveAt(index int)
//	Append(item any)
//	Prepend(item any)
//	Get(index int)
//}

type Node struct {
	value int
	next  *Node
}

type SingleLinkedList struct {
	head *Node
	len  int
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{}
}

func (s *SingleLinkedList) Append(val int) {
	n := Node{}
	n.value = val

	if s.len == 0 {
		s.head = &n
		s.len++
	} else {
		ptr := s.head
		for i := 0; i < s.len; i++ {
			if ptr.next == nil {
				ptr.next = &n
				s.len++
				break
			}
			ptr = ptr.next
		}
	}
}

func (s *SingleLinkedList) AppendAt(val int, index int) {

	var oldNode *Node
	newNode := Node{value: val}

	if index < 0 {
		return
	}

	if index == 0 {
		oldNode = s.head
		s.head = &newNode
		newNode.next = oldNode
	}

	if index > s.len {
		return
	}

	ptr := s.head
	for i := 0; i < s.len; i++ {
		if i == index-1 {
			oldNode = ptr.next
			ptr.next = &newNode
			newNode.next = oldNode
			s.len++
			return
		}
		ptr = ptr.next
	}
}

func (s *SingleLinkedList) Remove() {
	ptr := s.head
	for i := 0; i < s.len; i++ {
		if ptr.next.next == nil {
			ptr.next = nil
			s.len--
			break
		}
		ptr = ptr.next
	}
}

func (s *SingleLinkedList) Length() int {
	return s.len
}

func (s *SingleLinkedList) Print() {
	var printString string
	if s.len > 0 {
		ptr := s.head
		for i := 0; i < s.len; i++ {
			printString += fmt.Sprintf("%d ", ptr.value)
			ptr = ptr.next
		}
	}
	fmt.Println(printString)
}
