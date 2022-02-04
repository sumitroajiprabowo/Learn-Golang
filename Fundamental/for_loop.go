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

	slice := []string{"Danu", "Budi", "Caca"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// for with range
	for i, value := range slice {
		fmt.Println(i, "=", value)
	}

	// create map
	person := make(map[string]string)

	person["name"] = "Danu"
	person["age"] = "20"
	person["address"] = "Jakarta"

	// looping map
	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}
