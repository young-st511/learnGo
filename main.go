package main

import (
	"fmt"

	"github.com/young-st511/learnGo/myDict"
)

func main() {
	dictionary := myDict.Dictionary{}
	baseWord := "hello"

	dictionary.Add(baseWord, "Greeting")
	word, _ := dictionary.Search(baseWord)

	fmt.Println(word)

	dictionary.Delete("hello")
	word2, err := dictionary.Search(baseWord)

	if err == nil {
		fmt.Println(word2)
	} else {
		fmt.Println(err)
	}

}
