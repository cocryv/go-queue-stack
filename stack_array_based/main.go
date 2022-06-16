package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return 0
	}
	toRemove := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return toRemove
}

func (s Stack) isEmpty() bool {
	if len(s.items) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var myStack Stack

	// Checking if my stack is empty (true)
	fmt.Println(myStack.isEmpty())

	// Adding some items to my stack

	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	myStack.Push(4)

	// Checking if my stack is empty (false)
	fmt.Println(myStack.isEmpty())

	// Printing the items in the stack
	fmt.Println(myStack.items)

	// Removing the last item of the stack and return it
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())

	// Printing the items in the stack
	fmt.Println(myStack.items)

}
