package main

import "fmt"

func main() {
	fmt.Println("We are trying arrays from golang.")

	var fruits [4]string
	fruits[0] = "mango"
	fruits[1] = "guava"
	fruits[3] = "grapes"
	fmt.Println("Fruits are..", fruits)
	fmt.Println(len(fruits))

	var rollnos = [5]int{1, 2, 3}
	fmt.Println("Roll numbers are..", rollnos)
	fmt.Println(len(rollnos))
}
