package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"name"`
	Age int `json:"age,omitempty"`
	EmailAdress string `json:"email,omitempty"`
	Address Address `json:"address,omitempty"`
}

type Address struct {
	City string `json:"city"`
	State string `json:"state"`
}


func main() {

	person := Person{FirstName: "John", Age:30, EmailAdress: "fakeemail@gmail.com"}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error in Marshalling data:", err)
		return
	}
	fmt.Println(string(jsonData))

	person = Person{FirstName: "Smith", EmailAdress: "fakeemail2@gmail.com"}

	jsonData, err = json.Marshal(person)
	if err != nil {
		fmt.Println("Error in Marshalling data:", err)
		return
	}
	fmt.Println(string(jsonData))


	person1 := Person{
		FirstName: "Jane",
		Age: 21,
		EmailAdress: "jane@fakeemail.com",
		Address: Address{
			City: "New York",
			State: "NY",
		},
	}

	jsonData1, err := json.Marshal(person1)
	if err!=nil {
		fmt.Println("Error marshalling data in JSON:", err)
		return
	}
	fmt.Println(string(jsonData1))


	jsonDataEmployee := `{"full_name": "John Doe", "age": 26, "emp_id": "000874", "address": { "city": "New York", "state": "NY"}}`
	var employeeFromJson Employee
	err = json.Unmarshal([]byte(jsonDataEmployee), &employeeFromJson)
	if err != nil {
		fmt.Println("Error in Unmarshalling:", err)
		return
	}
	fmt.Println(employeeFromJson)
	fmt.Println("Age increade by 5:",employeeFromJson.Age+5)
	fmt.Println("Lives in:", employeeFromJson.Address.City)
	
	fmt.Println("------------------------------")


	listOfCityState := []Address {
		{City: "New York", State: "NY"},
		{City: "San Jose", State: "CA"},
		{City: "Las Vegas", State: "NV"},
		{City: "Modestro", State: "CA"},
		{City: "Clearwater", State: "FL"},
	}
	fmt.Println(listOfCityState)

	jsonData4, err := json.Marshal(listOfCityState)
	if err != nil {
		fmt.Println("error in marshalling the list:", err)
		return
	}
	fmt.Println(string(jsonData4))
	

	fmt.Println("-------------------------------")

	// Handling unknown json structures
	jsonData3 := `{"name": "John Doe", "age": 26, "address": { "city": "New York", "state": "NY"}}`
	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonData3), &data)
	if err != nil {
		fmt.Println("Error Unmarshalling JSON:", err)
		return
	}
	fmt.Println(data)
	fmt.Println(data["address"])
	fmt.Println(data["age"])


}

type Employee struct {
	FullName string `json:"full_name"`
	Age int `json:"age"`
	EmpID string `json:"emp_id"`
	Address Address `json:"address"`
}

