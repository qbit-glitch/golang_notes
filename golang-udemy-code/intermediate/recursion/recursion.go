package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
	fmt.Println(factorial(10))
	fmt.Println(sumOfDigits(12))
	fmt.Println(sumOfDigits(123))
	fmt.Println(sumOfDigits(125432))
}

func factorial(n int) int {

	// Base case: factorial or 0 is 1
	if n == 0 {
		return 1
	}
	// Recursive Case: Factorial of n is `n * factorial(n-1)`
	return n * factorial(n-1)
}

func sumOfDigits(n int) int{
	// Base Case
	if n<10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}