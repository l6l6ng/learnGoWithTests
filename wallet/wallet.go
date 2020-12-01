package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Println("address of balance in Deposit is ", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
var InsufficientFundsError = "cannot withdraw, insufficient funds"
func (w *Wallet) WithDraw(amount Bitcoin) error{

	if amount > w.balance {
		return errors.New(InsufficientFundsError)
	}

	w.balance -= amount
	return nil
}
