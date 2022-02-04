package main

import "fmt"

func getFullName(firstName string, lastName string) (string, string) {
	return firstName, lastName
}

func getNamaPanjang() (string, string, string) {
	return "Dani", "Saputri", "Jaya"
}

func main() {
	firstName, lastName := getFullName("Danu", "Saputra")
	fmt.Println(firstName, lastName)

	namaDepan, namaTengah, namaBelakang := getNamaPanjang()
	fmt.Println(namaDepan, namaTengah, namaBelakang)
}
