package main

import "fmt"

func main() {
	fmt.Println("Implementing stack using slices.")
	newStack := stack{-1, []int{}}
	newStack.push(2)
	newStack.push(-1)
	top, isemptyt := newStack.get_top()
	if isemptyt {
		fmt.Println("Stack is empty.")
	} else {
		fmt.Println("top is:", top)
	}
	newStack.push(7)
	pop, isempty := newStack.pop()
	if isempty {
		fmt.Println("Stack is empty, nothing to pop.")
	} else {
		fmt.Println("popped item is:", pop)
	}
	pop, isempty = newStack.pop()
	if isempty {
		fmt.Println("Stack is empty, nothing to pop.")
	} else {
		fmt.Println("popped item is:", pop)
	}
	pop, isempty = newStack.pop()
	if isempty {
		fmt.Println("Stack is empty, nothing to pop.")
	} else {
		fmt.Println("popped item is:", pop)
	}
	pop, isempty = newStack.pop()
	if isempty {
		fmt.Println("Stack is empty, nothing to pop.")
	} else {
		fmt.Println("popped item is:", pop)
	}

}

type stack struct {
	top int
	st  []int
}

func (s *stack) push(val int) {
	s.top++
	s.st = append(s.st, val)
}

func (s *stack) pop() (int, bool) {
	if s.top == -1 {
		return -1, true
	}
	s.top--
	val := s.st[len(s.st)-1]
	s.st = s.st[:len(s.st)-1]
	return val, false
}

func (s *stack) get_top() (int, bool) {
	if s.top == -1 {
		return -1, true
	}
	return s.st[s.top], false
}
