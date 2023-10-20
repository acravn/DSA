package linkedList

import "fmt"

type Stack struct {
	head *Node
	Len  int
}

func (s *Stack) Add(item int) {
	newNode := Node{value: item, next: s.head}
	s.head = &newNode
	s.Len++
}

func (s *Stack) Pop() (int, error) {
	if s.head == nil {
		return 0, fmt.Errorf("Empty stack\n")
	}

	value := s.head.value
	s.head = s.head.next
	return value, nil
}

func (s *Stack) Peek() int {
	return s.head.value
}
