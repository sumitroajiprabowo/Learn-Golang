package main

import "fmt"

func main() {
	name := "budi sudarsono"

	switch name {
	case "danu":
		fmt.Println("Ini Danu")
	case "Budi":
		fmt.Println("Ini Budi")
	case "erwin":
		fmt.Println("Ini Erwin")
	default:
		fmt.Println("Ini Default")
	}

	// Using short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Ini Gas")
	case false:
		fmt.Println("Ini Tidak Gas")
	}

	// switch without expression
	length := len(name)
	switch {
	case length > 5 && length < 10:
		fmt.Println("Lebih dari 5 Kurang dari 10")
	case length > 10:
		fmt.Println("Lebih dari 10")
	default:
		fmt.Println("Default")
	}

}
