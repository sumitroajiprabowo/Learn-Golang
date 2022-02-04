package main

import (
	"container/ring"
)

func main() {
	data := ring.New(3)
	// data.Value = 1
	// data.Next().Value = 2
	// data.Next().Next().Value = 3

	for i := 0; i < data.Len(); i++ {
		data.Value = i
		data = data.Next()
	}

	data.Do(func(value interface{})){
		fmt.Println(value)
	}

}
