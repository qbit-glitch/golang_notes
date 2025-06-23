package main

import (
	"errors"
	"fmt"
)

func main() {
	q, r := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)


	q, r = divideAgain(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", q, r)

	result, err := compare(3, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

}

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// function which returns an error as one of the values in the multiple return values
func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to Compare which is greater")
	}
}

func divideAgain(a, b int) (quotient int, remainder int){
	quotient = a/b
	remainder = a%b
	return
}
