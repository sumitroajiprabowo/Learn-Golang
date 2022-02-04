package main

import (
	"errors"
	"fmt"
)

func pembagi(nilai int, pembagi int) (int, error) {
	if nilai == 0 {
		return 0, errors.New("Nilai tidak boleh 0")
	}
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	}
	return nilai / pembagi, nil
}

func main() {
	var contohError error = errors.New("Ups Error")
	fmt.Println(contohError)

	// Example Error
	hasil, err := pembagi(10, 0) // 10 / 0
	if err != nil {
		fmt.Println(err) // Ups Error
	} else {
		fmt.Println(hasil) // 10
	}
}
