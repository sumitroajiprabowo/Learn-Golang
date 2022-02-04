package main

import "fmt"

func main() {

	// Create manual array

	// Array type string max 3 value
	var names [3]string

	// Add value by index array
	names[0] = "Budi"
	names[1] = "Gunawan"
	names[2] = "Haha"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Array langsung
	var values = [4]int{
		90,
		80,
		85,
		75,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// Melihat jumlah panjang di dalam array
	fmt.Println(len(values))
	fmt.Println(len(names))

}
