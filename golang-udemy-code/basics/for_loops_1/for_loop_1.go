package main

import "fmt"

func main() {

	// Simple iteration over a range
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// iterate over collections
	numbers := []int{2, 3, 4, 5, 6, 7, 8}
	for index, value := range numbers {
		fmt.Printf("Index: %d | Value: %d\n", index, value)
	}

	// Use of continue and break statements
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i == 5 {
			break
		}
	}

	// Asterisk Layout -> Christmas tree
	var rows int = 5
	for i := 1; i <= rows; i++ {
		// inner loop for blank spaces
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}

		// inner loop for asterisks
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	// range based for loops in go
	for i:= range 10 {
		fmt.Println(i)
	}

	fmt.Println("We have a lift off")
}
