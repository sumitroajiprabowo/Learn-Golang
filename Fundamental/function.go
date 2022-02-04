package main

import "fmt"

func helloWord() {
	fmt.Println("Hello World")
}

func main() {
	for i := 0; i < 10; i++ {
		helloWord()
	}
}
