// This file provides a concurrency-safe bank with one account.
package main

import "fmt"

type withdrawal struct {
	amount  int
	success chan bool
}

var deposits = make(chan int)   // send amount to deposit
var balances = make(chan int)   // receive balance
var wds = make(chan withdrawal) // remove amount from bank

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	suc := make(chan bool)
	wds <- withdrawal{amount, suc}
	return <-suc
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case w := <-wds:
			if w.amount > balance {
				w.success <- false
				continue
			}
			balance -= w.amount
			w.success <- true
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

func main() {

	fmt.Println("init balance:", Balance())

	fmt.Println("deposiit 50")
	Deposit(50)

	fmt.Println("new balance:", Balance())

	fmt.Println("deposiit 100")
	Deposit(100)

	fmt.Println("new balance:", Balance())

	fmt.Println("withdrawal 75:", Withdraw(75))

	fmt.Println("new balance:", Balance())

	fmt.Println("deposiit 10")
	Deposit(10)

	fmt.Println("new balance:", Balance())

	fmt.Println("withdrawal 100:", Withdraw(100))

	fmt.Println("new balance:", Balance())
}
