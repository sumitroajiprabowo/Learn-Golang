package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke", counter)
		counter++
	}

	// for with stement
	// init stement // condition stement // post stement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke", counter)
	}
}
