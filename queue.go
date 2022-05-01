package main

import "fmt"

type node struct {
	data int
	next *node
}

type queue struct {
	front *node
	rear  *node
}

func (q *queue) enqueue(x int) {
	newNode := node{x, nil}
	if q.front == nil {
		q.front = &newNode
		q.rear = q.front
	} else {
		q.rear.next = &newNode
		q.rear = q.rear.next
	}
}

func (q *queue) dequeue() {
	if q.front == nil {
		panic("Empty queue")
	}
	q.front = q.front.next

}

func (q *queue) printQueue() {
	temp := q.front
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func (q *queue) isEmpty() bool {
	return q.front == nil
}

func main() {
	q := queue{}
	q.enqueue(1)
	q.enqueue(2)
	q.dequeue()
	q.dequeue()
	q.enqueue(3)
	q.enqueue(4)
	q.printQueue()
}
