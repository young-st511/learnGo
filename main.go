package main

import "fmt"

func main() {
	//# Map - key, value의 쌍으로 이루어진 자료형
	young := map[string]string{"name": "young", "age": "24"}
	fmt.Println(young)

	// Iterate Map -> 이렇게 순환도 가능하다!
	for key, value := range young {
		fmt.Println(key, value)
	}
}
