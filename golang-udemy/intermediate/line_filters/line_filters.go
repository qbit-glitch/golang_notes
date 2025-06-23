package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("example.txt")
	if err!= nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer func(){
		fmt.Println("Closing file")
		file.Close()
	}()

	scanner := bufio.NewScanner(file)

	// Keyword
	keyword := "important"
	lineNumber := 1
	// Read and filter lines
	for scanner.Scan(){
		line := scanner.Text()
		if strings.Contains(line, keyword){
			updatedLine := strings.ReplaceAll(line, keyword, "necessary")
			fmt.Printf("%d Filtered Line: %v\n", lineNumber, line)
			fmt.Printf("%d Updated Line: %v\n", lineNumber, updatedLine)
			lineNumber ++

		}
	}

	err = scanner.Err()
	if err!= nil {
		fmt.Println("error scanning file:", err)
	}


}