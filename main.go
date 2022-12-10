package main

import "fmt"

func main() {
	// 상수 선언 법
	const NAME string = "youngHun"
	// 변수 선언 법
	var name string = "young"
	// 위와 같음 - 타입 추론을 해준다!!
	name2 := "young" // string

	fmt.Println("Hello World!" + name + name2)
}
