package main

import (
	"fmt"
	"regexp"
)

func main() {

	var regex *regexp.Regexp = regexp.MustCompile("[a-z]+") // compile regex

	fmt.Println(regex.MatchString("Hello World"))
	fmt.Println(regex.MatchString("Hello World!"))
	fmt.Println(regex.MatchString("Hello World?"))
	fmt.Println(regex.MatchString("Hello World."))
	fmt.Println(regex.MatchString("1234")) // false

	var text string = "Belajar Go" // "Belajar Go"
	fmt.Println(regex.FindAllString(text, -1))

}
