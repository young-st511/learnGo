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

# Bank Project

## Structure Methods

```go
/////////////////////////
// ./banking/banking.go
/////////////////////////

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
- `log.Fatalln(err)` - error 출력 후 프로그램 종료

```go
////////////////////////
// ./banking/banking.go
////////////////////////

// ...전략

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
}
```

```go
////////////////////////
// ./main.go
////////////////////////

func main() {
	// account := banking.Account{Owner: "young", Balance: 1000}
	account := banking.NewAccount("young")
	fmt.Println(account)
	account.Deposit(1000)
	fmt.Println(account.Balance())

	//# Go의 Error Handling(no try-catch, no exception)
	err := account.Withdraw(430)
	if err != nil {
		fmt.Println(err)
	}

	//# String method가 존재하는 경우 String meghod를 자동으로 호출시켜준다!
	fmt.Println(account)
}
```

- String method가 존재하는 경우 String meghod를 자동으로 호출시켜준다!

---

# Dictionary Project

```go
////////////////////////
// ./myDict/myDict.go
////////////////////////

package myDict

import "errors"

// Type alias
type Money int
// var money Money = 2

// just 'alias(가명)', structure 아님!! - key, value가 string
type Dictionary map[string]string

var errNotFound = errors.New("can not found the word")

//# 모든 type은 receiver를 이용하여 methods를 가질 수 있다!!

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	//# Map은 해당 key에 value가 존재하는지의 여부를 반환해준다!!!
	value, exists := d[word]

	if exists {
		return value, nil
	}

	return "", errNotFound
}
```

```go
package main

import (
	"fmt"

	"github.com/young-st511/learnGo/myDict"
)

func main() {
	dictionary := myDict.Dictionary{"first": "First word"}

	//# Go의 에러 핸들링!!!
	definition, err := dictionary.Search("second")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	fmt.Println(dictionary.Search("first"))
}
```

- Map은 해당 key에 value가 존재하는지의 여부를 반환해준다!!!
  - `value, exists := d[word]`

```go
////////////////////////
// ./banking/banking.go
////////////////////////

// Add a word to the dictionary
func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		fmt.Println("exist")
		return errWordExist
	}

	return nil
}
```

- error에 따른 switch 문으로 처리할 수도 있다!

### Update, Delete methods

```go
////////////////////////
// ./banking/banking.go
////////////////////////

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		// map docs 참고
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}

	return nil
}
```

```go
package main

import (
	"fmt"

	"github.com/young-st511/learnGo/myDict"
)

func main() {
	dictionary := myDict.Dictionary{}
	baseWord := "hello"

	dictionary.Add(baseWord, "Greeting")
	word, _ := dictionary.Search(baseWord)

	fmt.Println(word)

	dictionary.Delete("hello")
	word2, err := dictionary.Search(baseWord)

	if err == nil {
		fmt.Println(word2)
	} else {
		fmt.Println(err)
	}

}
```

---

# URL Checker Project

## map

- make 함수를 사용하지 않거나 map을 선언할 때 초기화 하지 않으면 map은 빈 map이 아닌 nil이 되어버린다!

```go
// var results = map[string]string -> panic error!! ❌
var results = map[string]string{}
// var results = make(map[string]string) -> ✅
```

## GoRoutine

- go 루틴을 동시에 실행하려고 할 때, 작업이 끝날 때까지 메인 함수는 기다려주지 않는다!

```go
func count(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, "hello!", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go count("young")
	go count("hun")
	time.Sleep(time.Second * 5) // i = 0 ~ 4
}
```

- 위 프로그램은 for루프를 5번까지만 돌고 끝나버림!

### Channels

- GoRoutine에서 데이터를 return 받기 위해서는 단순 return 사용이 불가능하다.
  - Channel을 열어줘야한다!!
- Channel에 값을 넣어주기 위해서는 다음과 같이 해야한다.
  - 화살표는 무조건 왼쪽 방향! `<-`
  - 송신: `ch <- true`
    - `c <- person + " Hi!"` 다른 자료형도 가능
  - 수신: `result := <- ch`
    - `fmt.Println(<-c)` 이렇게도 가능
- 런타임 시에 채널로부터 수신하는 코드가 있다면 데이터가 들어올 때까지 기다려준다!
  - `await` 랑 비슷한 듯!

```go
func count(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " Hi!"
}

func main() {
	c := make(chan string)
	people := [2]string{"young", "hun"}

	for _, person := range people {
		go count(person, c)
	}

	fmt.Println("Waiting for messages")
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("END")

	//# 이렇게 루프도 가능!!!!
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}
```

- 인수에서 chan 타입 지정 시에 send only, receive only를 화살표로 지정할 수 있다.
- send only: `chan<- ch`
- receive only: `<-chan ch`

### Fast!! URL Checker

```go
type result struct {
	url    string
	status string
}

func main() {
	ch := make(chan result)
	var results = map[string]string{}

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://www.naver.com/",
		"https://www.daum.net/",
	}

	for _, url := range urls {
		go hitURL(url, ch)
	}

	for range urls {
		res := <-ch
		results[res.url] = res.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, ch chan<- result) {
	fmt.Println("checking:", url)
	res, err := http.Get(url)
	status := "OK"

	if err != nil || res.StatusCode >= 400 {
		status = "FAILED"
	}

	ch <- result{url: url, status: status}
}
```

---

# Go Web Scrapper

## GoQuery

jQuery와 문법이 거의 비슷하며, Golang에서 DOM 요소를 조작할 수 있게 해준다!

## GoRoutine

```bash
# before
[Done] exited with code=0 in 10.419 seconds

# after
[Done] exited with code=0 in 2.874 seconds
```

GoRoutine과 channel을 이용하면 놀라운 퍼포먼스를 만들 수 있다...!

# Go Jobs Project

## Go Echo

- Go를 이용하여 간편하게 백엔드를 구성할 수 있다.
- 공식 문서가 잘 되어 있다.

## Go Buffalo

- 라우팅, 라이브 서버, 템플릿 등 백엔드 구성을 위한 많은 기능들을 제공
- Go의 Django 같은 것이라고 한다.
