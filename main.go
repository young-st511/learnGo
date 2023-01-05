package main

import (
	"fmt"

	"github.com/young-st511/learnGo/myDict"
)

func main() {
	dictionary := myDict.Dictionary{"first": "First word"}

	err2 := dictionary.Add("hello", "Greeting")

	if err2 != nil {
		fmt.Println(err2)
	}
	hello, _ := dictionary.Search("hello")

	fmt.Println(hello)

}
