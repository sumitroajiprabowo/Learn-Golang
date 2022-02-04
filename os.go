package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args
	fmt.Println("Arguments: ", args)
	fmt.Println(args)

	os.Args = []string{"say", "hello", "world"}
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])

	// os hostname
	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname", name)
	} else {
		fmt.Println("Error", err.Error())
	}

	// get from env
	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")
	fmt.Println("Username: ", username)
	fmt.Println("Password: ", password)

}
