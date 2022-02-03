package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"Novmber",
		"Desember",
	}
	var days = [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}
	var slice1 = months[4:7]

	// menampilkan slice
	fmt.Println(slice1) // [Mei Juni Juli]

	// mendapatkan panjang slice
	fmt.Println(len(slice1)) // 3

	// mendapatkan capacity
	fmt.Println(cap(slice1)) // 8

	// merubah nilai array maka slice juga berubah
	months[5] = "diubah"
	// Menampilkan nilai dari slice1 dengan array yang telah diubah
	fmt.Println(slice1) // [Mei diubah Juli]

	// menambahkan nilai pada slice maka array juga berubah
	slice1[2] = "Julio"
	// Menampilkan nilai dari array dengan slice yang telah diubah
	fmt.Println(months) // [Januari Februari Maret April Mei diubah Julio Agustus September Oktober Novmber Desember]

	daySlice1 := days[5:]

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"

	// [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]
	fmt.Println(days) // menampilkan array dengan nilai yang telah diubah

	// Append menambahkan
	daySlice2 := append(daySlice1, "Hari Libur")
	daySlice2[0] = "Ups"

	// Menampilkan slice dengan nilai yang telah ditambah
	fmt.Println(daySlice2) // [Ups Minggu Baru Hari Libur]

	// Menampilkan array days
	fmt.Println(days) // [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]
	// var daySlice1 = days[5:]
	// fmt.Println(daySlice1)

	// Membuat Slice dengan len 2 cap 5
	newSlice := make([]string, 2, 5)
	// Menambahkan nilai pada slice index 0
	newSlice[0] = "Hari"
	newSlice[1] = "Libur"

	fmt.Println(newSlice)      // [Hari Libur Baru]
	fmt.Println(len(newSlice)) // 2
	fmt.Println(cap(newSlice)) // 5

	// Membuat slice dengan nama copySlice dengan nilai dari newSlice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	// Copy Slice dari newSlice ke copySlice
	copy(copySlice, newSlice)
	fmt.Println(copySlice) // [Hari Libur Baru]

	// Perbedaan Aray dan Slice
	iniArrayBernilai := [4]int{1, 2, 3, 4}
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArrayBernilai) // [1 2 3 4]
	fmt.Println(iniArray)         // [1 2 3 4 5]
	fmt.Println(iniSlice)         // [1 2 3 4 5]

}
