package main

import (
	"flag"
	"fmt"
)

func main() {

	var host *string = flag.String("host", "localhost", "Put your host")
	var port *int = flag.Int("port", 3306, "Put your port")
	var user *string = flag.String("user", "root", "Put your user")
	var password *string = flag.String("password", "password", "Put your password")

	flag.Parse()

	fmt.Println("Host: ", *host)
	fmt.Println("Port: ", *port)
	fmt.Println("User: ", *user)
	fmt.Println("Password: ", *password)

}
