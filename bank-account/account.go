//Package account simulates a bank account with transactions.
package account

import (
	"sync"
)

const testVersion = 1

//Account is the basic structure for this package.
type Account struct {
	balance int64
	closed  bool
	sync.Mutex
}

//Open initializes the account, sets balance.
func Open(dep int64) *Account {
	if dep < 0 {
		return nil
	}
	return &Account{dep, false, sync.Mutex{}}
}

//Close closes the account and returns final balance.
func (acct *Account) Close() (int64, bool) {
	defer acct.Unlock()
	acct.Lock()
	if !acct.closed {
		acct.closed = true
		return acct.balance, true
	}
	return 0, false
}

//Balance returns current account balance.
func (acct *Account) Balance() (int64, bool) {
	return acct.Deposit(0)
}

//Deposit adds or subtracts money from the account.
func (acct *Account) Deposit(amt int64) (int64, bool) {
	defer acct.Unlock()
	acct.Lock()
	if acct.closed {
		return 0, false
	}
	if amt < 0 && acct.balance+amt < 0 {
		return acct.balance, false
	}
	acct.balance += amt
	return acct.balance, true
}
