package main

import "fmt"

var (
	sema    = make(chan struct{}, 20)
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // acquire token
	balance = balance + amount
	<-sema // release token
}

func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}

func main() {
	Deposit(1000)
	Deposit(1000)
	Deposit(1000)
	fmt.Println(Balance())
}
