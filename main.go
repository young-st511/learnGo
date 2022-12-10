package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	// for-in, for-of랑 비슷하다!
	// range는 index와 변환되는 값을 반환한다
	for index, number := range numbers {
		fmt.Println(index, number)
	}

	total := 0

	// 언더바 사용하면 컴파일러가 무시한다
	for _, number := range numbers {
		total += number
	}

	// 물론 전통적인 for문도 가능!!
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	return total
}

func main() {
	superAdd(11, 22, 33, 44, 55, 66, 77, 88)
}
