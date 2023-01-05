package main

import "fmt"

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
