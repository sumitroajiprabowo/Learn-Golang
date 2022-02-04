package main

import (
	"fmt"
	"reflect"
)

// Create a struct Sample
type Sample struct {
	Name string `required:"true" max:"10"` // tag
}

// Create a struct Example
type Example struct {
	Name string `required:"true" max:"10"` // tag
	Age  int    `required:"true"`          // required: true
}

// IsValid is a function to check if the struct is valid using tag
func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)           // get type of data
	for i := 0; i < t.NumField(); i++ { // loop through fields
		field := t.Field(i)                      // get field
		if field.Tag.Get("required") == "true" { // check if field is required
			value := reflect.ValueOf(data).Field(i).Interface() // get value of field
			if value == "" {                                    // check if value is empty
				return false // return false if value is empty
			}
		}
	}
	return true // return true if all field is valid
}

func main() {

	// create a struct Sample
	sample := Sample{Name: "Belajar Go"} // create struct Sample

	var sampleType reflect.Type = reflect.TypeOf(sample) // get type of struct Sample

	fmt.Println(sampleType.Name())                       // print name of struct Sample
	fmt.Println(sampleType.NumField())                   // print number of fields of struct Sample
	fmt.Println(sampleType.Kind())                       // print kind of struct Sample
	fmt.Println(sampleType.Field(0).Name)                // print name of field 0 of struct Sample
	fmt.Println(sampleType.Field(0).Tag)                 // print tag of field 0 of struct Sample
	fmt.Println(sampleType.Field(0).Tag.Get("required")) // print tag of field 0 of struct Sample
	fmt.Println(sampleType.Field(0).Tag.Get("max"))      // print tag of field 0 of struct Sample
	fmt.Println(IsValid(sample))                         // print if struct Sample is valid

	// create a struct Example
	examples := Example{Name: "Belajar Go", Age: 30} // create struct Example

	fmt.Println(IsValid(examples)) // print if struct Example is valid

}
