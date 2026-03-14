package banking

import (
	"errors"
	"fmt"
)

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

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
}
