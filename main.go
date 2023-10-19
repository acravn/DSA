package main

import (
	"DSA/datastructure/linkedList"
	"fmt"
)

func main() {
	fmt.Println("Hello DSA world")

	ll := linkedList.NewSingleLinkedList()

	ll.Append(1)
	ll.Append(3)
	ll.Append(4)
	ll.Append(5)
	ll.Print()

	// test AppendAt
	ll.AppendAt(2, 1)
	ll.Print()

	// test Remove
	ll.Remove()
	fmt.Println(ll.Length())
	ll.Print()
}
