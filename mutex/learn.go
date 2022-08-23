// What is a mutex ?
// Allow us to prevent a concurrent process
// from entering ciritcal portion of code
package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func deposit(value int, wq *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Deposit %d to accouunt with balance %d\n", value, balance)
	balance += value
	mutex.Unlock()
	wq.Done()

}
func withdraw(value int, wq *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Withdraw %d from account with balance %d\n", value, balance)
	
	mutex.Unlock()
	wq.Done()

}

func main() {
	fmt.Println("Hello world")
	balance = 1000
	var wq sync.WaitGroup
	wq.Add(2)
	go withdraw(700, &wq)
	go deposit(500, &wq)
	wq.Wait()
	fmt.Printf("New Balance %d", balance)

}
