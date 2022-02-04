package main

import (
	"belajar-go/helper"
	"fmt"
)

func main() {
	helper.SayHello("Budi")

	// cannot use sayGoodbye because it's private
	// helper.sayGoodbye("Gunawan") //error

	//example access variabel from package helper
	fmt.Println(helper.Application)

	// cannot using variable version because it's private
	// fmt.Println(helper.version) // error because version is private
}
