package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	// 일반적인 if statement
	// if age < 18 {

	// if문 선언과 동시에 변수 생성이 가능하다! - variable expression
	if koreanAge := age + 2; koreanAge < 20 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(18))
}
