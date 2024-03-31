package main

import "fmt"

func main() {
	fmt.Println("Exploring functions in golang!!")
	greeting()
	fmt.Println(add2int(4, 6))
	//soln, str := addnint(2,3,4,5)
	//fmt.Println(soln, str)
	fmt.Println(addnint(2, 3, 4, 5))
}

func greeting() {
	fmt.Println("Welcome!!")
}

func add2int(v1 int, v2 int) int {
	return v1 + v2
}

//Pro Functions
func addnint(values ...int) (int, string) {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum, "This is a Pro Function."
}
