package main

import "fmt"

func helloTo(name string, age int) {
	fmt.Println("Hello", name, "You are", age)
}

func main() {
	name := "Danu"
	helloTo(name, 20)
	helloTo("Budi", 20)
	helloTo("Andi", 21)
}
