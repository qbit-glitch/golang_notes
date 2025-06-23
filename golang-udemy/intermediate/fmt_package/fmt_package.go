package main

import "fmt"

func main() {

	s := fmt.Sprint("Hello", "World", 234, 456)
	fmt.Print(s)
	fmt.Print(s)

	s = fmt.Sprintln("Hello", "World", 234, 897)
	fmt.Print(s)
	fmt.Print(s)

	name := "John"
	age := 25
	sf := fmt.Sprintf("Name: %s, Age: %d\n", name, age)
	fmt.Print(sf)
	fmt.Print(sf)

	// Scanning Functions

	var nameInput string
	var ageInput int

	// fmt.Println("Enter your name and age: ")
	// fmt.Scan(&nameInput, &ageInput)
	// fmt.Printf("Name: %s, Age: %d\n", nameInput, ageInput)

	// fmt.Println("Enter your name and age: ")
	// fmt.Scanln(&nameInput, &ageInput)
	// fmt.Printf("Name: %s, Age: %d\n", nameInput, ageInput)
	
	fmt.Println("Enter your name and age: ")
	fmt.Scanf("%s - %d",&nameInput, &ageInput)
	fmt.Printf("Name: %s, Age: %d\n", nameInput, ageInput)
	
	err := checkAge(15)
	if err!= nil {
		fmt.Println("Error:", err)
	}

}


func checkAge(age int) error{
	if age<18 {
		return fmt.Errorf("Age %d is too young to drive.", age)
	}
	return nil
}