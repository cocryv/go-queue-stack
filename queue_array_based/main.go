package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Push(i int) {
	q.items = append(q.items, i)
}

func (q *Queue) Pop() int {
	if len(q.items) == 0 {
		return 0
	}
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {

	myQueue := Queue{}

	myQueue.Push(1)
	myQueue.Push(2)
	myQueue.Push(3)

	fmt.Println(myQueue.items)

	fmt.Println(myQueue.Pop())
	fmt.Println(myQueue.Pop())
	fmt.Println(myQueue.Pop())

	fmt.Println(myQueue.items)

}
