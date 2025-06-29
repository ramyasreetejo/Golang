package main

import (
	"fmt"
	"sync"
)

// A deadlock occurs when a set of goroutines are waiting for each other, and none can move forward.
// For instance, a deadlock occurs when a goroutine attempts to receive a message from an empty channel
// and has no other goroutines active. The "Receiver" goroutine, in this case, never receives a message.

func main() {
	wg := &sync.WaitGroup{}
	myCh := make(chan int, 1) // buffered channels

	wg.Add(2)
	go func(myCh chan<- int, wg *sync.WaitGroup) { //send only channel
		//close(myCh)
		fmt.Println("First go routine to pass data to a channel.")
		myCh <- 1
		myCh <- 2
		// close(myCh)
		wg.Done()
	}(myCh, wg)

	go func(myCh <-chan int, wg *sync.WaitGroup) { //receive only channel
		fmt.Println("Second go routine to receive data from a channel.")
		//close(myCh)
		val, isChannelOpen := <-myCh
		fmt.Println(<-myCh)
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
