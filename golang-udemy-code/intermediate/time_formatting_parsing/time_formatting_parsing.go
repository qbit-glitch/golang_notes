package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("\n---- Time Formatting and Parsing ----\n")

	// Mon Jan 2 15:04:05 MST 2006

	layout := "2006-02-02T15:04:05Z07:00"
	str := "2024-07-04T14:30:18Z"

	t, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	fmt.Println(t)

	str1 := "Jul 03, 2024 03:18 PM"
	layout1 := "Jan 02, 2006 03:04 PM"
	t1, err := time.Parse(layout1, str1)
	fmt.Println(t1)


}