package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q *Queue) Enqueue(i int) {
	node := Node{value: i}

	if q.size == 0 {
		q.head = &node
		q.tail = &node
	} else {
		q.tail.next = &node
		q.tail = &node
	}
	q.size++

}

func (q *Queue) Dequeue() interface{} {
	if q.size == 0 {
		return nil
	} else {
		toRemove := q.head
		q.head = q.head.next
		q.size--

		return toRemove.value
	}
}

func main() {

	myQueue := Queue{}

	myQueue.Enqueue(1)
	myQueue.Enqueue(3)
	myQueue.Enqueue(2)

	a := myQueue.Dequeue()
	fmt.Println(a)

}
