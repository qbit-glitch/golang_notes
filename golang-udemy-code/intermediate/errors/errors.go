package main

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error){
	if x<0 {
		return 0, errors.New("math error: square root of a negative number")
	}
	// Compute the square root
	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Error: Empty data")
	}
	// Process the data
	return nil
}





func main() {

	result1, err1 := sqrt(-16)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println(result1)
	}
	

	data := []byte{}
	if err := process(data); err != nil {
		fmt.Println(err)
		
	} else {
		fmt.Println("Data Processed Successfully")
	}
	
	// error interface of built-in package
	if err1 := eprocess(); err1 != nil {
		fmt.Println(err1)
	}
	
	if err2 := readData(); err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Data read successfully.")
	}

	

}



type myError struct{
	message string
}

func (m *myError) Error() string{
	return fmt.Sprintf("Error: %s", m.message)
}

func eprocess() error {
	return &myError{"Custom Error Message"}
}


func readConfig() error{
	return errors.New("config error")
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}