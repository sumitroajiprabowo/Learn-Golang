package main

import "fmt"

func main() {
	// Deklarasi Konstanta
	// Nilai Tetap ( Tidak Bisa diubah )
	// Universal Value
	const firstName = "Danu"
	const lastName = "Budi"
	const age = 20

	// Tidak error jika nilai const tidak di panggil
	// fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)

	// Multiple Const
	const (
		namaDepan    = "Danu"
		namaBelakang = "Budi"
		usia         = 20
	)
	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
	fmt.Println(usia)

}
