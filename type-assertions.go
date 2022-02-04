package main

import "fmt"

func random() interface{} {
	return "OK"
}

func example() interface{} {
	return true
}

func main() {
	var result interface{} = random()
	var resultString string = result.(string)
	fmt.Print(resultString)

	// var resultInt int = result.(int) // panic
	// fmt.Print(resultInt)

	var hasil interface{} = example()
	switch value := hasil.(type) {
	case bool:
		fmt.Println("Value", value, "is a bool")
	case string:
		fmt.Println("Value", value, "is a string")
	case int:
		fmt.Println("Value", value, "is a int")
	default:
		fmt.Println("Unknown type")
	}
}
