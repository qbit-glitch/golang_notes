package main

import "fmt"

func main() {

	day:= "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday")
	case "Saturday", "Sunday" :
		fmt.Println("Its a weekend")
	default:
		fmt.Println("Invalid day")
	}

	num := 2
	switch {
	case num > 1:
		fmt.Println("Greater than 1")
	case num ==2 :
		fmt.Println("Number is 2")
	default :
		fmt.Println("Less than 1")
	}

	checkType(10)
	checkType(3.18)
	checkType("Hello")
	checkType(true)

}

func checkType(x interface{}){
	switch x.(type){
	case int:
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("Its a float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown type")
	}
}