# learnGo
Golang을 연습합니다

# Golang

# 😇 기초

## export

```go
// PascalCase로 쓰여진 변수나 함수는 export 된다!
func main() {
	// ❌ error!
	fmt.println("Hellow World!")

	// ✅ OK!
	fmt.Println("Hellow World!")
}
```

## 변수 선언 및 타입 선언

```go
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
```

## 함수

### 함수 타입 선언

```go
package main

import "fmt"

// x int, y int
func multiply(x, y int) int {
	return x * y
}

func main() {
	fmt.Println(multiply(2, 3))
}
```

### 다중 반환 함수 (Multiple results)

```go
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
```

### 스프레드 연산자(?)

```go
// 스프레드연산자(...) 를 이용하면 여러 개를 받을 수 있다!
func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	repeatMe("abc", "young", "nico", "dal", "jung")
}
```

## 반복문
```go
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
```

