package main

import (
	"fmt"

	"github.com/ongdv/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("OngDV")
	account.Deposit(10)
	fmt.Println(account)
}
