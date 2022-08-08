package main

import (
	"fmt"
	"strings"
)

// name := "OngDV" // short form Only working on function
// var name = "OngDV"

// func multiply(a int, b int) int {
// 	return a * b
// }

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	// const
	// const name string = "Ongdv"
	// name = "ongdv94" // is not working

	// variable
	// name := "OngDV" //var name string = "OngDV"
	// name = "ongdv94"

	// fmt.Println(multiply(2, 2))
	// totalLength, upperName := lenAndUpper("ongdv")
	// totalLength, _ := lenAndUpper("ongdv") // _ is ignore value
	// fmt.Println(totalLength, upperName)
	repeatMe("ongdv", "kyeongwoo", "oh", "doubleO")

}
