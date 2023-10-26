package linkedList

import "fmt"

type Queue struct {
	head *Node
	tail *Node
	Len  int
}

func (q *Queue) enqueue(item any) {
	newNode := Node{value: item}
	// assumption queue is empty
	if q.head == nil {
		q.head = &newNode
		q.Len++
		return
	}
	if q.tail == nil {
		q.tail = &newNode
		q.head.next = q.tail
		q.Len++
		return
	}
	q.tail.next = &newNode
	q.tail = &newNode
	q.Len++
}

func (q *Queue) deque() (any, error) {
	if q.head == nil {
		return 0, fmt.Errorf("Empty queue\n")
	}
	nextNode := q.head.next
	val := q.head.value
	q.head = nextNode
	q.Len--
	return val, nil
}

func (q *Queue) peek() any {
	return q.head.value
}
