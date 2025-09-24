package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(time.Duration(5) * time.Second)
	fmt.Println(x)
}

// read write mutex
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

	time.Sleep(time.Duration(5) * time.Second)
	fmt.Println("Total Balance:", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Locking...", user1.Name)
	user1.Change(amount * -1)

	time.Sleep(time.Duration(1) * time.Second)

	user2.Lock()
	fmt.Println("Locking...", user2.Name)
	user2.Change(amount)

	time.Sleep(time.Duration(1) * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := &UserBalance{
		Name:    "John",
		Balance: 100000000,
	}

	user2 := &UserBalance{
		Name:    "Deep",
		Balance: 1000000000,
	}

	fmt.Println(user1)
	fmt.Println(user2)

	go Transfer(user1, user2, 2000)
	go Transfer(user2, user1, 20000)

	time.Sleep(time.Duration(5) * time.Second)
	fmt.Println(user1)
	fmt.Println(user2)
}
