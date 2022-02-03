package main

import "fmt"

func main() {
	name := "Danu"

	if name == "Budi" {
		fmt.Println("Benar")
	} else if name == "Danu" {
		fmt.Println("Hallo")
	} else {
		fmt.Println("Salah")
	}

	// Using if short statement
	if length := len(name); length > 5 {
		fmt.Println("Gas")
	} else {
		fmt.Println("Tidak Gas")
	}
}
