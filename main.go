package main

import (
	"fmt"
)

func main() {
	a := 2
	b := a
	a = 10 // a = 10, b = 2
	fmt.Println(a, b)

	aa := 2
	bb := &aa
	//# & -> 메모리 주소를 볼 수 있다
	//# * -> 해당 주소에 들어가 값을 볼 수 있다
	aa = 6
	fmt.Println(aa, *bb) // aa = 6, *bb = 6
	*bb = 0
	fmt.Println(aa, *bb) // aa = 0, *bb = 0

}
