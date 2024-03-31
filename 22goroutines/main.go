package main

import (
	"fmt"
	"time"
)

func main() {
	go greeter("hello")
	greeter("World")
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}
