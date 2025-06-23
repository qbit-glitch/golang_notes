package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("\n---- bufio Package ---- ")

	readerExample()
	writerExample()
}

func readerExample() {

	reader := bufio.NewReader(strings.NewReader("Hello World! bufio package tutorial\n"))

	// Reading the byte slice
	data := make([]byte, 20)
	n, err := reader.Read(data)
	if err != nil {
		fmt.Println("error reading the string:", err)
		return
	}
	fmt.Printf("Read %d bytes: %s\n", n, data[:n])

	// Reading the string with delimeters
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error eading the string:", err)
	}
	fmt.Println("Read String:", line)
	fmt.Println()
}

func writerExample() {

	writer := bufio.NewWriter(os.Stdout)
	fmt.Println("---- Writer ----")
	// Writing byte slice
	data := []byte("Hello, bufio package!\n")
	n, err := writer.Write(data)
	if err != nil {
		fmt.Println("error writing:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)

	// Flush the buffer to ensure all data is written to os.Stdout

	err = writer.Flush()
	if err != nil {
		fmt.Println("error flushing writer:", err)
		return
	}

	// Writing Strings
	str := "This is a string\n"
	n, err = writer.WriteString(str)
	if err!= nil {
		fmt.Println("error writing string:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)

	// Flush the buffer
	err = writer.Flush()
	if err != nil {
		fmt.Println("error flushing writer:", err)
	}



}
