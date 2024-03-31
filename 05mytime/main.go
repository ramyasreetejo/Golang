package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to my time study of golang")

	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2001, time.June, 3, 10, 30, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}
