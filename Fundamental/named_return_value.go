package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Danu"
	middleName = "Saputra"
	lastName = "Jaya"

	// return
	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName, middleName, lastName)
}
