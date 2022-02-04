package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	}
	return map[string]string{
		"name": name,
	}

}

func main() {
	var person map[string]string = nil
	fmt.Println(person)

	data := newMap("")
	if data == nil {
		fmt.Println("Data is nil")
	} else {
		fmt.Println(data)
	}
}
