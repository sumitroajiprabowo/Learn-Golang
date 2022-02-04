package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Budi",
		"address": "Jl. Kebon Jeruk",
		"age":     "20",
	}

	person["hobby"] = "Coding"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["age"])
	fmt.Println(person["hobby"])

	// Mengambil jumlah data dari map
	fmt.Println(len(person))

	// Membuat maps baru
	var book map[string]string = make(map[string]string)

	//Menambahkan data ke dalam map dengan key dan value
	book["title"] = "Golang"
	book["author"] = "Budi"
	book["publisher"] = "Golang.org"

	// Menampilkan data dari map
	fmt.Println(book)

	// Menghapus data dari map
	delete(book, "title")

	fmt.Println(book)
}
