package account

import (
	"math"
	"sync"
)

type Account struct {
	balance int64
	open    bool
	mu      sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	return &Account{balance: amount, open: true, mu: sync.Mutex{}}
}

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	if !a.open {
		a.mu.Unlock()
		return 0, false
	}

	a.mu.Unlock()
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	if !a.open {
		a.mu.Unlock()
		return 0, false
	}

	if amount < 0 && math.Abs(float64(amount)) > float64(a.balance) {
		a.mu.Unlock()
		return 0, false
	}

	a.balance += amount
	a.mu.Unlock()
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	if !a.open {
		a.mu.Unlock()
		return 0, false
	}

	payout := a.balance
	a.balance = 0
	a.open = false
	a.mu.Unlock()
	return payout, true
}
