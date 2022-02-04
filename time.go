package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now() // get current time
	fmt.Println(now)  // print current time

	fmt.Println(now.Year())       // print current year
	fmt.Println(now.Month())      // print current month
	fmt.Println(now.Day())        // print current day
	fmt.Println(now.Hour())       // print current hour
	fmt.Println(now.Minute())     // print current minute
	fmt.Println(now.Nanosecond()) // print current nanosecond

	// Formatting time
	fmt.Println(now.Format("2006-01-02 15:04:05")) // print current time with format
	fmt.Println(now.Format("2006-01-02"))          // print current time with format
	fmt.Println(now.Format("15:04:05"))            // print current time with format

	utc := time.Date(2022, time.February, 2, 15, 4, 5, 0, time.UTC) // create time with UTC
	fmt.Println(utc.Local())                                        // print local time

	layout := "2006-01-02"                       // layout for parsing
	parse, _ := time.Parse(layout, "2022-02-01") // parse time
	fmt.Println(parse)                           // print parsed time

}
