package database

import "fmt"

var connection string

func init() {
	fmt.Println("Initializing database connection") // this will be executed only once
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
