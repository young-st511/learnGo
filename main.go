package main

import (
	"fmt"

	"github.com/young-st511/learnGo/banking"
)

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
