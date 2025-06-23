package main

import (
	"fmt"
	"os"
)


func checkError(err error){
	if err!= nil {
		panic(err)
	}
}


func main() {

	templFileName := "temporaryFile"
	tempFile, err := os.CreateTemp("", templFileName)
	checkError(err)

	fmt.Println("Temporary File created:", tempFile.Name())

	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	tempDir, err := os.MkdirTemp("","GoCourseTempDir")
	checkError(err)

	defer os.RemoveAll(tempDir)
	
	fmt.Println("Temporary directory created:", tempDir)
}