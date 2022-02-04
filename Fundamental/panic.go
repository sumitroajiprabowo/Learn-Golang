package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error bos")
	}
	// jika error maka ini tidak akan dijalankan
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
}
