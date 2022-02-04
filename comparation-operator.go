package main

import "fmt"

func main() {
	var name1 = "Budi"
	var name2 = "Budi"

	var result bool = name1 == name2 // with boolean type

	var reresult = name1 == name2 // without boolean type

	fmt.Println(result)
	fmt.Println(reresult)

	var value1 = 100
	var value2 = 200
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 >= value2)
	fmt.Println(value1 <= value2)
	fmt.Println(value1 != value2)
}
