package main

import (
	"fmt"
)

// defer follows LIFO
func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Hello")
	MyDefer()

}

func MyDefer() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
