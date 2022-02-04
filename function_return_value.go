package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello World"
	}
	return "Hello " + name
}

func main() {
	result := getHello("Danu")
	fmt.Println(result)
}
