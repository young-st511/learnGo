package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	// switch문도 variable expression 사용 가능!
	switch koreanAge := age + 2; koreanAge {
	case 18:
		return true
	case 10:
		return false
	case 50:
		return true
	}
	return false
}

func canIDrink2(age int) bool {
	// switch문에 조건문도 사용 가능
	switch koreanAge := age + 2; {
	case koreanAge < 18:
		return false
	case koreanAge > 70:
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(18))
}
