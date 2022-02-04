package main

import "fmt"

// customer is a struct
type customer struct {
	Name, Address string
	Age           int
}

// changeValueAge is a function that change value of age
func changeValueAge(c *customer) {
	c.Age = 30 // change value of age
}

// changeValueAddress is a function that change value of address
func changeValueAddress(c *customer) {
	c.Address = "Pemalang"
}

func main() {
	customer1 := customer{"Danu", "Jakarta", 20} // struct literal
	// pass by value (copy)
	customer2 := customer1 // copy

	customer2.Name = "Budi" // change value of customer2

	fmt.Println(customer2) // {Budi Jakarta 20}
	fmt.Println(customer1) // {Danu Jakarta 20}

	// pass by reference (pointer)
	customer3 := &customer1 // pointer to customer1
	customer3.Name = "Caca" // change value of customer3

	fmt.Println(customer3) // &{Caca Jakarta 20}

	customer3 = &customer{"Putri", "Bandung", 20} // pointer to customer

	fmt.Println(customer3) // &{Putri Bandung 20}

	// pass by value (copy)
	customer4 := &customer1                       // pointer to customer1
	*customer4 = customer{"Putra", "Bandung", 20} // change value of customer4

	// all value of customer4 will be changed
	fmt.Println(customer1) // {Putra Bandung 20}
	fmt.Println(customer2) // {Putra Bandung 20}
	fmt.Println(customer3) // &{Putra Bandung 20}
	fmt.Println(customer4) // &{Putra Bandung 20}

	// function new
	customer5 := new(customer)    // create new customer
	customer5.Name = "Danu"       // set value of customer5
	customer5.Address = "Jakarta" // set value of customer5
	customer5.Age = 20            //
	fmt.Println(customer5)        // &{Danu Jakarta 20}

	customer6 := customer{"Putri", "Bandung", 20} // struct literal
	changeValueAge(&customer6)                    // pass by reference
	fmt.Println(customer6)                        // {Putri Bandung 30}

	customer7 := customer6         // copy
	changeValueAddress(&customer7) // pass by reference
	fmt.Println(customer7)         // {Putri Pemalang 20}

}
