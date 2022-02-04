package main

import (
	"fmt"
	"strings"
)

func main() {

	// Mengecek string yang yang ada di dalam flag
	// exampe using strings.Contains()
	fmt.Println(strings.Contains("Hello World", "World")) // true
	fmt.Println(strings.Contains("Hello World", "Hey"))   // false

	// example using split()
	// memisahkan string menjadi array
	fmt.Println(strings.Split("Hello World", " ")) // ["Hello" "World"]

	// example using strings.ToUpper()
	// mengubah string menjadi huruf kapital
	fmt.Println(strings.ToUpper("Hello World")) // HELLO WORLD

	// example using strings.ToLower()
	// mengubah string menjadi huruf kecil
	fmt.Println(strings.ToLower("Hello World")) // hello world

	// example using strings.TrimSpace()
	// menghilangkan spasi di awal dan akhir string
	fmt.Println(strings.TrimSpace(" Hello World ")) // Hello World

	// example using strings.Trim()
	// menghilangkan spasi di awal dan akhir string
	fmt.Println(strings.Trim("   Hello World   ", " "))

	// example using strings.TrimLeft()
	// menghilangkan spasi di awal string
	fmt.Println(strings.TrimLeft("   Hello World   ", " ")) // Hello World

	// example using strings.TrimRight()
	// menghilangkan spasi di akhir string
	fmt.Println(strings.TrimRight("   Hello World   ", " ")) //   Hello World

	// example using strings.Replace()
	// mengganti string
	fmt.Println(strings.Replace("Hello World", "World", "Dolly", -1)) // Hello Dolly

	// example using strings.Join()
	// menggabungkan string menjadi 1 string
	fmt.Println(strings.Join([]string{"Hello", "World"}, " ")) // Hello World

	// example using strings.Index()
	// mencari index string
	fmt.Println(strings.Index("Hello World", "World")) // 6

	// example using strings.LastIndex()
	// mencari index string
	fmt.Println(strings.LastIndex("Hello World", "World")) // 6

	// example using strings.Count()
	// menghitung jumlah string
	fmt.Println(strings.Count("Hello World", "l")) // 2

	// example using strings.ContainsAny()
	// mengecek apakah string berisi karakter tertentu
	fmt.Println(strings.ContainsAny("Hello World", "lo")) // true

	// example using strings.ContainsRune()
	// mengecek apakah string berisi karakter tertentu
	fmt.Println(strings.ContainsRune("Hello World", 'l')) // true

	// example using strings.Fields()
	// mengecek apakah string berisi karakter tertentu
	fmt.Println(strings.Fields("Hello World")) // ["Hello" "World"]

}
