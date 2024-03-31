package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	rwmut := &sync.RWMutex{}

	arr := []int{}

	wg.Add(4)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		mut.Lock()
		fmt.Println("This is the first R.")
		arr = append(arr, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		mut.Lock()
		fmt.Println("This is the second R.")
		arr = append(arr, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		mut.Lock()
		fmt.Println("This is the third R.")
		arr = append(arr, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, rwmut *sync.RWMutex) {
		rwmut.RLock()
		fmt.Println("This is the fourth R.")
		fmt.Println(arr)
		rwmut.RUnlock()
		wg.Done()
	}(wg, rwmut)

	wg.Wait()
	fmt.Println(arr)
}
