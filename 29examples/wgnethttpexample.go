package main

import (
	"fmt"
	"net/http"
	"sync"
)

func getUrls(urls *[]string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range len(*urls) {
		var weburl string = (*urls)[i]
		fmt.Println(weburl)
		resp, err := http.Get(weburl)
		if err != nil {
			fmt.Printf("Error getting %v err: %v\n", weburl, err)
		} else {
			fmt.Printf("Status code: %v\n", resp.StatusCode)
		}
	}
}

func getUrl(url string, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	resp, err := http.Get(url)
	var op string
	if err != nil {
		op = fmt.Sprintf("Error getting %v err: %v\n", url, err)
	} else {
		op = fmt.Sprintf("Status code of %v: %v\n", url, resp.StatusCode)
	}
	ch <- op
}

func main() {
	urls := []string{
		"https://www.instagram.com/",
		"https://github.com/ramyasreetejo",
		"https://www.linkedin.com/in/ramya-sree-tejomurtula/"}
	wg := &sync.WaitGroup{}
	ch := make(chan string, 3)
	// wg.Add(1)
	// go getUrls(&urls, wg)
	for i := range urls {
		wg.Add(1)
		go getUrl(urls[i], wg, ch)
	}
	wg.Wait()
	close(ch)
	for op := range ch {
		fmt.Println(op)
	}
	// time.Sleep(time.Second)
}
