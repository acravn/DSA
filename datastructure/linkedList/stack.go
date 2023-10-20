package linkedList

import "fmt"

type Stack struct {
	head *Node
	Len  int
}

func (s *Stack) Add(item any) {
	newNode := Node{value: item, next: s.head}
	s.head = &newNode
	s.Len++
}

func (s *Stack) Pop() (any, error) {
	if s.head == nil {
		return 0, fmt.Errorf("Empty stack\n")
	}

	value := s.head.value
	s.head = s.head.next
	s.Len--
	return value, nil
}

func (s *Stack) Peek() any {
	return s.head.value
}
