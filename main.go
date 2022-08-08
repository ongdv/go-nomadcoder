package main

import (
	"fmt"

	"github.com/ongdv/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("OngDV")
	fmt.Println(account)
}
