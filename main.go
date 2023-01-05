package main

import (
	"fmt"
	"log"

	"github.com/young-st511/learnGo/banking"
)

func main() {
	// account := banking.Account{Owner: "young", Balance: 1000}
	account := banking.NewAccount("young")
	fmt.Println(account)
	account.Deposit(1000)
	fmt.Println(account.Balance())

	//# Go의 Error Handling(no try-catch, no exception)
	err := account.Withdraw(2000)
	if err != nil {
		// 출력과 함께 프로그램 종료
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}
