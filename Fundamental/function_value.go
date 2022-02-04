package main

import "fmt"

// Function value
func thankYou(name string) string {
	return "Thank you " + name
}

func main() {

	// Function value
	fmt.Println(thankYou("Danu"))

	// Function value
	hasil := thankYou("Danu")
	fmt.Println(hasil)

	// or use function value
	sayThankYou := thankYou
	result := sayThankYou("Danu")
	fmt.Println(result)
}
