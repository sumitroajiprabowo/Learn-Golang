package main

import "fmt"

func main() {

	// Variabel dengan Tipe Data String
	var name string

	// Nilai Variabel name
	name = "Belajar"
	// Menampilkan Variabel name
	fmt.Println(name)

	// Reusable Variabel name
	name = "Golang"
	// Menampilkan Variabel name
	fmt.Println(name)

	// Variabel String dengan value
	var friendName = "Budi"
	fmt.Println(friendName)

	// Variabel Integer dengan Value
	var umur = 20
	fmt.Println(umur)

	// Inisialisasi Variabel dengan :=
	country := "Indonesia"
	fmt.Println(country)

	// Reusable Variabel

	country = "Malaysia"
	fmt.Println(country)

	// Multiple Variabel
	var (
		firstName = "Danu"
		lastName  = "Budi"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(firstName + lastName)

}
