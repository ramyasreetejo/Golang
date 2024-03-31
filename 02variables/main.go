package main

import (
	"fmt"
)

// loginTime := 03:30 (not correct)
const Logintoken = "dhjksmsjkwsa" //this is a public variable

func main() {
	var name string = "Ramyaaa"
	fmt.Println(name)
	fmt.Printf("The type of the variable is: %T. \n", name)

	//we cant access characters in a string using str[i] etc and cant swap also, so we convert string
	//to rune which is char array so that we can access individual characters and also swap them, later
	//can convert back using string() func... for demo see 18permutations file code.

	var num int
	fmt.Println(num)
	fmt.Printf("The type of the variable is: %T. \n", num)

	//implicit way w out mentioning type
	var isSafe = true
	fmt.Println(isSafe)

	//no var style or short declaration operator for declaration+assignment
	marks := 98.5
	fmt.Println(marks)
	fmt.Printf("The type of the variable is: %T. \n", marks)

	fmt.Println(Logintoken)
	fmt.Printf("The type of the variable is: %T. \n", Logintoken)
}
