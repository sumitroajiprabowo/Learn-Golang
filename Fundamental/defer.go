package main

import "fmt"

func logging() {
	fmt.Println("Aplikasi Selesai")
}

func runAplication() {
	defer logging()
	fmt.Println("Aplikasi Berjalan")
}

func exampleError(value int) {
	defer logging()
	fmt.Println("Aplikasi Berjalan Tapi Nanti Error")
	result := value / value
	fmt.Println("Result", result)
}

func main() {
	runAplication()
	exampleError(0)
}
