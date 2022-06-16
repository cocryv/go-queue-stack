package main

import (
	"fmt"
	"strings"
)

type Node struct {
	value int
	next  *Node
}

func (n *Node) toString() string {
	if n != nil {
		return fmt.Sprint(n.value)
	} else {
		return ""
	}

}

type Stack struct {
	head *Node
	size int
}

func (s Stack) toString() string {
	sb := strings.Builder{}
	for iterator := s.head; iterator != nil; iterator = iterator.next {
		sb.WriteString(iterator.toString() + "\t")
	}
	return sb.String()
}

func (s *Stack) Push(i int) {
	node := Node{value: i}
	s.size++
	if s.head == nil {
		s.head = &node
	} else {
		node.next = s.head
		s.head = &node
	}
}

func (s *Stack) Pop() *Node {
	if s.head == nil {
		return nil
	} else {
		toRemove := s.head
		s.head = s.head.next
		s.size--
		return toRemove
	}
}

func main() {

	var myStack Stack

	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	myStack.Push(4)

	fmt.Println(myStack.toString())

	deleted := myStack.Pop()
	fmt.Println(deleted.toString())

	fmt.Println(myStack.toString())

}
