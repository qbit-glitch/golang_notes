package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("ouput1.txt")
	if err != nil {
		fmt.Println("error creating file:", err)
		return
	}
	defer file.Close()

	// Write data to a file
	data := []byte ("Hello, World!!\n\n\n\n\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("error writing data:", err)
		return
	}

	fmt.Println("Data has been written to file successfully.")

	file, err = os.Create("output2.txt")
	if err != nil {
		fmt.Println("error creating file:", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString("Hello, Go lang\n")
	if err != nil {
		fmt.Println("error wrting string to file:", err)
		return
	}
	fmt.Println("Writing to output2.txt complete.")


}