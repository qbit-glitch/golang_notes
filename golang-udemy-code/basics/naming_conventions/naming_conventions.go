package main

import "fmt"

type EmployeeGoogle struct{
	FirstName string
	LastName string
	Age int
}

func main(){
	/* Naming Conventions
	1. PascalCase -> structs, interfaces, enums
		eg: CalculateArea, UserInfo, NewHTTPRequest, etc

	2. snake_case
		eg: user_id, first_name, http_request
	
	3. UPPERCASE -> constants

	4. mixedCase
		eg: javaScript, htmlDocument, isValid	
	*/

	const MAXTRIES = 5

	var employeeID = 10001
	fmt.Println("Employee ID : ", employeeID)
}