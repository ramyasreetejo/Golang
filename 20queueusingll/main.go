package main

import "fmt"

func main() {
	fmt.Println("Implementing Queue w LL.")
	q := Queue{nil, nil, 0}
	q.enqueue(1)
	q.enqueue(2)
	fmt.Println(q.front.data)
	ele := q.dequeue()
	if ele != nil {
		fmt.Println(ele.data)
	} else {
		fmt.Println("Queue is empty.")
	}
	fmt.Println(q.front.data)
	ele = q.dequeue()
	if ele != nil {
		fmt.Println(ele.data)
	} else {
		fmt.Println("Queue is empty.")
	}
	ele = q.dequeue()
	if ele != nil {
		fmt.Println(ele.data)
	} else {
		fmt.Println("Queue is empty.")
	}
}

type Node struct {
	data int
	next *Node
}

type Queue struct {
	front *Node
	rare  *Node
	size  int
}

func (q *Queue) enqueue(val int) {
	q.size++
	newNode := &Node{val, nil}
	if q.front == nil && q.rare == nil {
		q.front = newNode
		q.rare = newNode
		return
	}
	q.rare.next = newNode
	q.rare = newNode
}

func (q *Queue) dequeue() *Node {
	if q.front == nil && q.rare == nil {
		return nil
	}
	q.size--
	temp := q.front
	if q.front.next == nil {
		q.front = nil
		q.rare = nil
	} else {
		q.front = q.front.next
	}
	return temp
}
