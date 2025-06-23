package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("\n--- NUMBER PARSING ---")

	numStr := "89"
	numVal, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Invalid string for parsing to integer")
	} else {
		fmt.Println("Parsed Integer :", numVal)
	}

	numStr = "50"
	numParse, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		fmt.Println("error parsing the string to integer")
	} else {
		fmt.Println("Parsed Integer :", numParse)
	}

	binStr := "1010001"
	bitInt, err := strconv.ParseInt(binStr, 2, 64)
	if err != nil {
		fmt.Println("error parsing the string to binary")
	} else {
		fmt.Println("Parsed Binary :", bitInt)
	}
	floatStr := "3fc"
	floatVal, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("error parsing the string to float")
	} else {
		fmt.Println("Parsed Float :", floatVal)

	}

	hexStr := "ff"
	hexInt, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("error parsing the string to hexadecimal")
	} else {
		fmt.Println("Parsed Hex :", hexInt)
	}

}
