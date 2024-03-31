package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers concept.")

	var ptr1 *int
	fmt.Println(ptr1)

	var num int = 8
	var ptr2 *int = &num // (or) var ptr2 = &num (same)
	fmt.Println(ptr2, "\n", *ptr2, "\n", *ptr2+2)
}
