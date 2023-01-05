package main

import "fmt"

func main() {
	//# Array
	// Go의 배열 선언법 - Array has length!
	names := [5]string{"jung", "young", "hun"}
	names[3] = "haha"
	names[4] = "lala"
	// names[5] = "kaka" // error!
	fmt.Println(names)

	//# Slice
	// 선언은 배열과 비슷하다
	names2 := []string{"jung", "young", "hun"}
	// 새로운 Slice를 return!
	names2 = append(names2, "lala")
	fmt.Println(names2)
}
