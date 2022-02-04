package main

import "fmt"

// customer struct
type customer struct {
	Name, Address string // Name, Address
	Age           int    // Age
	Actived       bool   // Actived
}

// myCustomer method using struct method and pointer
func (c *customer) myCustomer() {
	c.Name = "Mr. " + c.Name // change value of Name
}

func (c *customer) setActivated() {
	c.Actived = true // change value of Actived
}

func main() {
	danu := customer{"Danu", "Jakarta", 20, false} // create struct literal
	danu.myCustomer()                              // call method
	fmt.Println(danu.Name)                         // Mr. Danu

	danu.setActivated()       // call method
	fmt.Println(danu.Actived) // true
	fmt.Println(danu)         // {Mr. Danu Jakarta 20 true}
}
