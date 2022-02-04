package main

import "fmt"

func main() {
	name := "Danu"
	counter := 0

	// closure
	increment := func() {
		// counter := 1
		// replace value variabel name
		// name = "Budi"

		// use this
		name := "Caca"

		// becarefull
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
