package main

import (
	"container/list"
	"fmt"
)

func main() {
	// create a new list
	data := list.New() // data is a pointer to a list
	data.PushBack(1)   // add 1 to the list
	data.PushBack(2)   // add 2 to the list
	data.PushBack(3)   // add 3 to the list

	// iterate over the list
	fmt.Println(data.Front().Value) // 1

	// iterate over the list
	fmt.Println(data.Back().Value) // 3

	// iterate over the list
	// data.Front() is a pointer to the first element of the list
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// iterate over the list
	// data.Back() is a pointer to the last element of the list
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
