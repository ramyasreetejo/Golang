package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	myCh := make(chan int, 1) // buffered channels

	wg.Add(2)
	go func(myCh chan<- int, wg *sync.WaitGroup) { //send only channel
		//close(myCh)
		fmt.Println("First go routine to pass data to a channel.")
		myCh <- 1
		myCh <- 2
		close(myCh)
		wg.Done()
	}(myCh, wg)

	go func(myCh <-chan int, wg *sync.WaitGroup) { //receive only channel
		fmt.Println("Second go routine to receive data from a channel.")
		//close(myCh)
		val, isChannelOpen := <-myCh
		//fmt.Println(<-myCh)
		//fmt.Println(<-myCh)
		if isChannelOpen {
			fmt.Println(val)
		} else {
			fmt.Println("Channel is closed hence value returns 0")
		}
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
