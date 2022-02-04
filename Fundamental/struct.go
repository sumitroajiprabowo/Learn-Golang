package main

import "fmt"

type customer struct {
	Name, Address string
	Age           int
}

func main() {
	var danu customer
	danu.Name = "Danu"
	danu.Address = "Jakarta"
	danu.Age = 20

	fmt.Println(danu)
	fmt.Println(danu.Name)
	fmt.Println(danu.Address)
	fmt.Println(danu.Age)

	// Create Struct Literal
	budi := customer{
		Name:    "Budi",
		Age:     20,
		Address: "Jakarta",
	}
	fmt.Println(budi)
	fmt.Println(budi.Name)
	fmt.Println(budi.Age)
	fmt.Println(budi.Address)

	// Create Struct Literal with Short Hand
	caca := customer{"Caca", "Bandung", 20}
	fmt.Println(caca)
	fmt.Println(caca.Name)
	fmt.Println(caca.Age)
	fmt.Println(caca.Address)
}
