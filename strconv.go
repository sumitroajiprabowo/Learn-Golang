package main

import (
	"fmt"
	"strconv"
)

func main() {

	// example using ParseBool
	// string to bool
	boolean, err := strconv.ParseBool("true") // true
	if err != nil {
		fmt.Println("Error", err.Error()) // Error parsing bool
	} else {
		fmt.Println("Boolean: ", boolean) // true
	}

	// example of ParseInt
	// string to int
	number, err := strconv.ParseInt("123", 10, 64) // base, bitSize
	if err != nil {
		fmt.Println("Error", err.Error()) // Error: invalid syntax
	} else {
		fmt.Println("Number: ", number) // Number: 123
	}

	// FormatInt
	// int to string
	value := strconv.FormatInt(999, 10) // int64, base
	fmt.Println("Value: ", value)       // Value: 999

	valueInt, err := strconv.Atoi("123") // string to int
	if err != nil {
		fmt.Println("Error", err.Error()) // Error: invalid syntax
	} else {
		fmt.Println("Value: ", valueInt) // Values: 123
	}

	// example strconv.Itoa
	// int to string
	valueInt = 333
	value = strconv.Itoa(valueInt) // int to string
	fmt.Println("Value: ", value)  // Value: 333

}
