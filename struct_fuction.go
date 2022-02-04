package main

import "fmt"

type myCustomer struct {
	Name, Address string
	Age           int
}

func siHi(customer myCustomer, name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

// struct method
func (customer myCustomer) siHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var danu myCustomer
	danu.Name = "Danu"
	danu.Address = "Jakarta"
	danu.Age = 20

	siHi(danu, "Budi")

	// execute struct method
	danu.siHello("Caca")

	// fmt.Println(danu)
	// fmt.Println(danu.Name)
	// fmt.Println(danu.Address)
	// fmt.Println(danu.Age)

	// // Create Struct Literal
	// budi := customer{
	// 	Name:    "Budi",
	// 	Age:     20,
	// 	Address: "Jakarta",
	// }
	// fmt.Println(budi)
	// fmt.Println(budi.Name)
	// fmt.Println(budi.Age)
	// fmt.Println(budi.Address)

	// // Create Struct Literal with Short Hand
	// caca := customer{"Caca", "Bandung", 20}
	// fmt.Println(caca)
	// fmt.Println(caca.Name)
	// fmt.Println(caca.Age)
	// fmt.Println(caca.Address)
}
