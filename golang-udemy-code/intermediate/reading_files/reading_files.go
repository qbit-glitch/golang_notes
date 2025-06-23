package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("output1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func() {
		fmt.Println("Closing open file")
		file.Close()
	}()

	fmt.Println("File opened successfully!")
	
	// readFileEntirely(file)
	
	scanner := bufio.NewScanner(file)

	// Read line by line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}

func readFileEntirely(file *os.File) {
	// Read contents of the opened file
	data := make([]byte, 1024)
	_, err := file.Read(data)
	if err != nil {
		fmt.Println("error reading data from file:", err)
		return
	}
	fmt.Println("File Content:", string(data))

}
