package main

import "fmt"

func main() {

	// opration && Dan
	// operation || Atau
	// operation ! Kebalikan

	// Operation && and / dan
	// true && true = true
	// true && false = false
	// false && true = false
	// false && false = false

	// Operation || or / atau
	// true || true = true
	// true || false = true
	// false || true = true
	// false || false = false

	// Operation ! negation / Kembalikan
	// !true = false
	// !false = true

	var nilaiAkhir = 80
	var absensi = 75

	var lulusNilaiAkhir bool = nilaiAkhir >= 80
	var lulusAbsensi bool = absensi >= 80

	fmt.Println(lulusNilaiAkhir)
	fmt.Println(lulusAbsensi)

	var lulus = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)

	fmt.Println(nilaiAkhir >= 80 && absensi >= 80)

}
