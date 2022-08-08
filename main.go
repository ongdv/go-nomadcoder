package main

import (
	"fmt"
	"log"

	"github.com/ongdv/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("OngDV")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}
