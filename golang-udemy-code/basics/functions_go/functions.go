package main

import "fmt"

func main() {

	// func(){
	// 	fmt.Println("Hello, Anonymous Function")
	// }()

	greet := func(){
		fmt.Println("Hello, Anonymous Function")
	}
	greet()


	operations := add
	result2 := operations(2,4)
	fmt.Println("2 + 4 =", result2)

	// Passing a function as an argument
	result := applyOperation(5,3, add)
	fmt.Println("5 + 3 =", result)

	//  Returning and using a function
	multiplyBy2 := createMultiplier(2)
	fmt.Println("7 * 2 =", multiplyBy2(7))
	
}

func add(a,b int) int{
	return a+b
}

// Function that takes a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x,y)
}


// Function that returns a function
func createMultiplier(factor int) func(int) int{
	return func(x int) int{
		return x * factor
	}
}