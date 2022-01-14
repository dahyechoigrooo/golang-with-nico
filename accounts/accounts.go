package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var erroNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
// *, & : pointer를 사용하면 메모리를 적게 차지 하기 때문에 실행 속도가 빨라진다.
func NewAccount(owner string) *Account { // *Account : 위에 선언한 Account를 그대로 사용하기 위함
	account := Account{owner: owner, balance: 0}
	return &account // &account : 새로 생성하지 않고 기존에 존재하던 데이터를 그대로 return한다.
}

// Deposit x amount on your account (예금)
func (a *Account) Deposit(amount int) {
	//fmt.Println("Gonna deposit", amount)
	a.balance += amount
}

// Balance of your account (잔액)
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account (인출)
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return erroNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
