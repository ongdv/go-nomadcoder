package main

import "fmt"

func main() {
	names := []string{"ongdv", "bam", "ddong"}
	names2 := append(names, "12")
	fmt.Println(names, names2)
}
