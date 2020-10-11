package main

import "fmt"

// Queue
type Queue struct {
	items []int
}

//Enqueue
func (q *Queue) enqueue(i int) {
	q.items = append(q.items, i)
}

//Dequeue
func (q *Queue) dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.enqueue(100)
	myQueue.enqueue(200)
	myQueue.enqueue(300)
	fmt.Println(myQueue)
	myQueue.dequeue()
	fmt.Println(myQueue)
}
