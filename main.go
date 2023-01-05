package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("request failed")

func main() {
	//# make 함수를 사용하지 않거나 map을 선언할 때 초기화 하지 않으면 map은 빈 map이 아닌 nil이 되어버린다!
	// var results = map[string]string -> panic error!! ❌
	var results = map[string]string{}
	// var results = make(map[string]string) -> ✅

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
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("checking:", url)
	res, err := http.Get(url)

	if err != nil || res.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
