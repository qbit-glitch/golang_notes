package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Command:", os.Args[0])

	for i, arg := range os.Args {
		fmt.Println("Argument", i ,":", arg)
	}

	// Define flags
	var name string
	var age int
	var male bool

	flag.StringVar(&name, "n", "John", "name of the user")
	flag.IntVar(&age, "a", 18, "age of the user")
	flag.BoolVar(&male, "male", true, "gender of the user")

	flag.Parse()

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Male:", male)

}