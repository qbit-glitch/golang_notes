package main

import "fmt"

func main() {
	var age int = 25

	if age >= 50 {
		if age >= 65 {
			fmt.Println("Retired Citizen")
		} else {
			fmt.Println("Senior Citizen but not retired")
		}
	} else if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Not an Adult")
	}

	number := 11

	if number%2 == 0 && number%5 == 0 {
		fmt.Println("Divisible by 2 and 5")
	} else if number%2 == 0 || number%5 == 0 {
		fmt.Println("Divisible by either 2 or 5")
	} else {
		fmt.Println("Neither Divisible by 2 nor 5")
	}

}
