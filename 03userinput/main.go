package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	//bufio is like buffer and in it we store keyboard i/p using os.Stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of our pizza out of 5")

	//comma ok, comma err syntax, we can also say input, err if we are using the err to print or send somewhere
	//read the buffer until you find \n or end line.
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating us with a ", input)
	fmt.Printf("Type of input is %T.", input)
}
