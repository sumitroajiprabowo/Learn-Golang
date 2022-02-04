package main

import "fmt"

type animal interface {
	breathe()
	walk()
}

// Another Example Zero Interface
func ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}

}

func main() {
	var a animal
	fmt.Println(a)

	// Example Zero Interface
	var data interface{} = ups(2)
	fmt.Println(data)
}
