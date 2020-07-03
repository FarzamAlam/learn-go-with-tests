package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

var ErrOutOfLimit = errors.New("Not Enough Balance")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrOutOfLimit
	}
	w.balance = w.balance - amount
	return nil
}
