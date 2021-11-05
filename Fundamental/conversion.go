package main

import "fmt"

func main() {
	// Variabel Nilai int32
	var nilai32 int32 = 129
	// Konversi ke int64
	var nilai64 = int64(nilai32)

	//Konversi ke int8
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// Tipe Data String
	var firstName = "Danu"
	//byte alias uint8
	var e byte = firstName[0]
	// Konvert dari byte ke String
	var eString = string(e)

	// Menampilkan nilai variabel
	fmt.Println(firstName)
	fmt.Println(eString)

}
