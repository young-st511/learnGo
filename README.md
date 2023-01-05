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

## 조건문

```go
func canIDrink(age int) bool {
	// 일반적인 if statement
	// if age < 18 {

	// if문 선언과 동시에 변수 생성이 가능하다! - variable expression
	if koreanAge := age + 2; koreanAge < 20 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(18))
}
```

### switch문

```go
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
```

## Pointer

```go
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
```

## Array

### Array & Slice

```go
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
```

### Maps

```go
func main() {
	//# Map - key, value의 쌍으로 이루어진 자료형
	young := map[string]string{"name": "young", "age": "24"}
	fmt.Println(young)

	// Iterate Map -> 이렇게 순환도 가능하다!
	for key, value := range young {
		fmt.Println(key, value)
	}
}
```

### Structs

```go
//# Go에는 Class, 즉, constructor가 없다!!
//# Structure의 type 지정!
type person struct {
	name     string
	age      int
	favFoods []string
}

func main() {
	//# Structure(구조체) - JS의 Object와 비슷함!
	favFoods := []string{"banana", "pasta", "egg"}

	// 이렇게 선언 가능하지만 명확하지 않음
	young := person{"young", 21, favFoods}

	// Object처럼 key 명시 가능!!
	young = person{name: "young", age: 21, favFoods: favFoods}
	fmt.Println(young.age)
	fmt.Println(young)
}
```

---

# Bank & Dictionary Projects

## Structure Methods

```go
// ./banking/banking.go

package banking

import "errors"

//# Public을 위해 PascalCase를 사용한다!!!!
//# Account struct
// type Account struct {
// 	Owner   string
// 	Balance int
// }

// 하지만 Public은 보안에 문제가 많다
type Account struct {
	owner   string
	balance int
}

// NewAccount - Account의 참조를 반환함
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//# receiver convention - method를 만들 struct의 첫 글자를 따서 소문자로!
//# pointer에 주의한다!!!!!!!!!!

// Deposit - Account의 method 생성
func /* pointer receiver - Account의 원본 포인터 */ (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your Account
func /* Account의 복사본 */ (a Account) Balance() int {
	return a.balance
}

var errNoMoney = errors.New("can't withdraw")

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		// return errors.New("can't withdraw, you are poor")
		return errNoMoney
	}

	a.balance -= amount
	return nil // error의 null
}
```

- method는 receiver라는 것을 선언하여 만들어준다.
  - **Go에는 Class와 construtor가 없다!! 직접 초기화 함수를 만들어야한다.**
  - receiver convention: method를 만들 struct의 첫 글자를 따서 소문자로!
- 내부 값을 직접 수정할 시 pointer에 주의한다!
  - **this 대신 pointer 사용!**
  - 포인터 사용하지 않으면 자동으로 복사본 생성
    - 오히려 안전해서 좋아!
-
