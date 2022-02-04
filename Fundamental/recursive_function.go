package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i // result = result * i
	}
	return result
}

func factorialRecursive(value int) int {
	// if value == 0 and you use code value < 1 == infinite loop
	if value <= 1 {
		return 1
	}
	return value * factorialRecursive(value-1)
}

func main() {
	loop := factorialLoop(5)
	fmt.Println(loop)
	fmt.Println(5 * 4 * 3 * 2 * 1)
	fmt.Println(factorialRecursive(5))

	factorial := factorialRecursive(5)
	fmt.Println(factorial)
}
