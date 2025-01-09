package account

import "sync"

// Define the Account type here.
type Account struct {
	open    bool
	balance int64
	mutex   sync.RWMutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount, mutex: sync.RWMutex{}, open: true}
}

func (a *Account) Balance() (int64, bool) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	return a.balance, a.open
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if a.open && a.balance+amount >= 0 {
		a.balance += amount
		return a.balance, true
	}
	return a.balance, false

}

func (a *Account) Close() (int64, bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if a.open {
		a.open = false
		previousBalance := a.balance
		a.balance = 0
		return previousBalance, true
	}
	return 0, false
}
