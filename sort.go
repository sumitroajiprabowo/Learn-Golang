package main

import (
	"fmt"
	"sort"
)

// create struct User
type User struct {
	Name string // Name is a string
	Age  int    // Age is an int
}

// UserSlice is a slice of User
type UserSlice []User

// Len is a method for sort.Interface
func (value UserSlice) Len() int {
	return len(value) // return length of slice
}

// Less is a method for sort.Interface
func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age // sort by age
}

// Swap is a method for sort.Interface
func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i] // swap two elements
}

func main() {

	// create slice of User
	users := []User{
		{"Danu", 30}, // create User
		{"Budi", 20}, // create User
		{"Caca", 10}, //
	}

	fmt.Println("Before sorting:") // print before sorting
	fmt.Println(users)             // print slice of User
	sort.Sort(UserSlice(users))    // sort slice of User
	fmt.Println("After sorting:")  // print after sorting
	fmt.Println(users)             // print slice of User
}
