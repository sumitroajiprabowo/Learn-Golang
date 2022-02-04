package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total

}

func main() {
	total := sumAll(10, 10, 30, 40, 50)
	fmt.Println(total)

	// slice to variadic function
	slice := []int{1, 2, 3, 4, 5}
	total = sumAll(slice...)
	fmt.Println(total)
}
