package main

import "fmt"

func main() {
	ongdv := map[string]string{"name": "ongdv", "age": "12"}

	for key, value := range ongdv {
		fmt.Println(key, value)
	}
	fmt.Println(ongdv)
}
