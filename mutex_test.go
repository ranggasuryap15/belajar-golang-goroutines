package belajar_golang_gouroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	var x = 0
	var mutex sync.Mutex

	for i := 1; i < 1000; i++ {
		go func() {
			for j := 1; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	// mutex itu untuk mengatasi race condition, ini agak sedikit lambat, tapi tidak terlalu lambat banget karena cuma beda beberapa nano second

	time.Sleep(7 * time.Second)
	fmt.Println("Counter :", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance :", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (userBalance *UserBalance) Lock() {
	userBalance.Mutex.Lock()
}

func (userBalance *UserBalance) Unlock() {
	userBalance.Mutex.Unlock()
}

func (userBalance *UserBalance) Change(amount int) {
	userBalance.Balance = userBalance.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	user1.Change(-amount)
	fmt.Println("Lock user 1", user1.Name)
	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user 2", user2.Name)
	user2.Change(amount)
	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Rangga",
		Balance: 1000,
	}

	user2 := UserBalance{
		Name:    "Kevin",
		Balance: 1000,
	}

	go Transfer(&user1, &user2, 100)
	//go Transfer(&user2, &user1, 100)

	time.Sleep(3 * time.Second)

	fmt.Println("User", user1.Name, ", Balance", user1.Balance)
	fmt.Println("User", user2.Name, ", Balance", user2.Balance)

}
