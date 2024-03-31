package main

import (
	"fmt"
	"net/http"
	"sync"
)

//wait groups and mutex demo

var wg sync.WaitGroup //pointer
var mut sync.Mutex    //pointer
var endpoints []string

func main() {
	urls := []string{
		"https://www.google.com/",
		"https://www.instagram.com/",
		"https://www.youtube.com/",
		"https://twitter.com/home",
		"https://www.linkedin.com/"}
	for i := range urls {
		go getStatusCode(urls[i])
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(endpoints)
}

func getStatusCode(web string) {
	defer wg.Done()
	res, err := http.Get(web)
	if err != nil {
		fmt.Println("OOPS some error while getting status code!!")
	} else {
		mut.Lock()
		endpoints = append(endpoints, web)
		fmt.Printf("The status code for %v is %v.\n", web, res.StatusCode)
		mut.Unlock()
	}
}
