package main

import "fmt"
func main() {

	// panic(interface)
	process(10)
	process(-3)
}

func process(input int){
	defer fmt.Println("Defered 1")
	defer fmt.Println("Defered 2")

	if input < 0{
		fmt.Println("Before Panic")
		panic("Input must be a non-negative number")
		
		// defer fmt.Println("Defered 3")
		
	}
	fmt.Println("Processing input:", input)
}