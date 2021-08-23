package concurrency

import (
	"fmt"
	"sync"
	//"time"
)

var (
	mutex sync.Mutex
	//mutex2 sync.Mutex // to produce dead lock
	balance int = 1000
)

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	//mutex2.Lock() // Reverse sequence of Mutex to produce dead lock
	fmt.Printf("Deposit: %v in Balance: %v\n", value, balance)
	balance += value;
	//mutex2.Unlock()
	mutex.Unlock()
	wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup) {
	//mutex2.Lock() // Reverse sequence of Mutex to produce dead lock
	//time.Sleep(time.Second) // will produce dead lock if withdraw routine executes first.
	mutex.Lock() 
	fmt.Printf("Withdraw: %v from Balance: %v\n", value, balance)
	balance -= value;
	mutex.Unlock()
	//mutex2.Unlock()
	wg.Done()
}

func MutexMain() {
	var wg sync.WaitGroup
	wg.Add(2)

	go deposit(500, &wg)
	go withdraw(700, &wg)

	wg.Wait()

	fmt.Println("Balance: ", balance)
}