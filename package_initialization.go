package main

import (
	"belajar-go/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
