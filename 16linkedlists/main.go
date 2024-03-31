package main

import "fmt"

func main() {
	fmt.Println("Implementing Linked lists in go!!")
	ll := LinkedList{nil, 0}
	ll.insertll(1)
	ll.insertll(2)
	ll.insertll(3)
	ll.printll()
	ll.deletell(5)
	ll.printll()
	ll.insertll(9)
	ll.printll()
	ll.deletell(9)
	ll.printll()
	ll.insertll(4)
	ll.insertll(5)
	ll.deletell(1)
	ll.printll()
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (ll *LinkedList) insertll(val int) {
	newNode := &Node{val, nil, nil}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	temp := ll.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
	newNode.prev = temp
	ll.size += 1
}

func (ll *LinkedList) deletell(val int) {
	if ll.head == nil {
		return
	}
	if ll.head.value == val {
		ll.size -= 1
		if ll.head.next == nil {
			ll.head = nil
			return
		}
		ll.head = ll.head.next
		ll.head.prev = nil
		return
	}
	temp := ll.head
	for temp.next != nil && temp.next.value != val {
		temp = temp.next
	}
	if temp.next == nil {
		return
	}
	temp.next = temp.next.next
	if temp.next != nil {
		temp.next.prev = temp
	}
	ll.size -= 1
}

func (ll *LinkedList) printll() {
	if ll.head == nil {
		return
	}
	temp := ll.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}
