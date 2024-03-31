package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Exploring if else and switch in golang!!")
	loginCount := 10

	if loginCount < 10 {
		fmt.Println("Login count is less than 10 so valid.")
	} else if loginCount > 10 {
		fmt.Println("Login count is more than 10 so invalid, try again after 5 min.")
	} else {
		fmt.Println("Caution: Login count is 10.")
	}

	num := 9
	if num%2 == 0 {
		fmt.Printf("%v is even.\n", num)
	} else {
		fmt.Printf("%v is odd.\n", num)
	}

	if num := 5; num == 5 {
		fmt.Println("The received number is 5.")
	} else {
		fmt.Println("The received number is not 5.")
	}

	//if err != nil{
	// do something
	//}

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("The dice number is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You can move 1 spot")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
		fallthrough
	case 4:
		fmt.Println("You can move 4 spot")
		fallthrough
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 6 spot and throw the dice again ")
	default:
		fmt.Println("What was that!!")
	}
}
