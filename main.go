package main

import (
	"fmt"
)

// 스프레드연산자..? 를 이용하면 여러 개를 받을 수 있다!
func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	repeatMe("abc", "young", "nico", "dal", "jung")
}
