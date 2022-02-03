package main

import "fmt"

func main() {

	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	// Augmented Assigment

	var i = 10
	var k = 10

	i += 10 // i = i + 10
	k -= 9  // i = i -9
	fmt.Println(i)
	fmt.Println(k)

	// Unary Operator

	i++
	fmt.Println(i)

	k++
	fmt.Println(k)

	var negatif = -100
	var positif = +100 //

	fmt.Println(negatif)
	fmt.Println(positif)

}
