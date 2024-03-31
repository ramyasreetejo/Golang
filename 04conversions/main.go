package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to conversions of user input"
	fmt.Println(welcome)

	//bufio is like buffer and in it we store keyboard i/p using os.Stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of our pizza out of 5")

	//comma ok, comma err syntax, we can also say input, err if we are using the err to handle it like send somewhere or to print etc.
	//read the buffer until you find \n or end line.
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)

	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("We added 1 to your rating to make it ", numrating+1)
	}
}
