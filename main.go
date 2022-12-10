package main

import "fmt"

// PascalCase로 쓰여진 변수나 함수는 export 된다!
func main() {
	// ❌ error!
	fmt.println("Hellow World!")

	// ✅ OK!
	fmt.Println("Hellow World!")
}
