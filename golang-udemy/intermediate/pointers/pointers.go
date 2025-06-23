package main

import "fmt"

func main() {

	fmt.Println("Pointers in Go !!")

	var ptr *int
	var a int = 10

	if ptr == nil {
		fmt.Println("Pointer is nil")
	}

	ptr = &a   // referencing

	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println(*ptr)    // deferencing a pointer

	modifyValue(ptr)
	fmt.Println(a)
}

func modifyValue(ptr *int){
	*ptr ++
}