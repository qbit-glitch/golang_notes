package main

import "fmt"

func main() {

	process(10)
}

func process(i int) {
	defer fmt.Println("Deferred i value:",i)	// since it is encountered earlier before i is incremented that's why it has the original value of i
	defer fmt.Println("First Deferred statement executed")
	defer fmt.Println("Second Deferred statement executed")
	defer fmt.Println("Third Deferred statement executed")

	i++
	fmt.Println("Normal execution statement")
	fmt.Println("Value of i:",i)
}