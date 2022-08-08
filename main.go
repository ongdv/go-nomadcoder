package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"sprits", "dekkila"}
	ongdv := person{name: "오경우", age: 29, favFood: favFood}
	fmt.Println(ongdv)
}
