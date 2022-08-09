package main

import (
	"fmt"

	"github.com/ongdv/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	sWord, err1 := dictionary.Search(baseWord)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(sWord)
	}
	deleteError := dictionary.Delete(baseWord)
	// deleteError := dictionary.Delete("baseWord")
	if deleteError != nil {
		fmt.Println(deleteError)
	}
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
