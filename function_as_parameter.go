package main

import "fmt"

type filtering func(string) string

// func type declaration
func sayWithFilter(name string, filter filtering) {
	result := filter(name)
	fmt.Println("Hello", result)
}

// Create function as parameter
// func sayWithFilter(name string, filter func(string) string) {
// 	result := filter(name)
// 	fmt.Println("Hello", result)
// }

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}
	// else
	return name
}

// Create function as parameter
func kataKotor(name string) string {
	// slice of string
	slice := []string{"Anjing", "Babi", "Bajingan"}
	for _, value := range slice {
		if name == value {
			return "***"
		}
	}
	return name
}

func main() {
	sayWithFilter("Danu", spamFilter)
	sayWithFilter("Anjing", spamFilter)
	sayWithFilter("Bajingan", kataKotor)
}
