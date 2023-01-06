package main

import (
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

func main() {
	ch := make(chan result)
	var results = map[string]string{}

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://www.naver.com/",
		"https://www.daum.net/",
	}

	for _, url := range urls {
		go hitURL(url, ch)
	}

	for range urls {
		res := <-ch
		results[res.url] = res.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, ch chan<- result) {
	fmt.Println("checking:", url)
	res, err := http.Get(url)
	status := "OK"

	if err != nil || res.StatusCode >= 400 {
		status = "FAILED"
	}

	ch <- result{url: url, status: status}
}
