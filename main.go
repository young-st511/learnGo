package main

import (
	"fmt"
	"strings"
)

// ,를 이용해서 여러 개를 반환할 수 있다!
// Go에는 여러 멋진 패키지가 존재한다
func lenAdnUpper(name string) (int, string) {
	return len(name), strings.ToUpper((name))
}

func main() {
	// 이렇게 불러올 수 있다!! _(underBar)를 사용하면 무시할 수 있다
	totalLeng, upperName := lenAdnUpper("young")
	fmt.Println(totalLeng, upperName)
}
