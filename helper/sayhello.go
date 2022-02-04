package helper

import "fmt"

var Application = "Belajar Golang"
var version = 1

// Bisa di akses dari luar packagae helper because first letter is capital
func SayHello(name string) {
	fmt.Println("Hello from package helper", name)
}

// Tidak bisa diakses dari luar package helper because first letter is lowercase
func sayGoodbye(name string) {
	fmt.Println("Goodbye from package helper", name)
}
