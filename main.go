package main

import (
	"fmt"

	"github.com/young-st511/learnGo/myDict"
)

func main() {
	dictionary := myDict.Dictionary{"first": "First word"}

	//# Go의 에러 핸들링!!!
	definition, err := dictionary.Search("second")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	fmt.Println(dictionary.Search("first"))
}
