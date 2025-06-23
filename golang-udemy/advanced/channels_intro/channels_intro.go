package main

import (
	"fmt"
	"time"
)

func main() {

	// variable = make(chan Type)
	greeting := make(chan string)
	greetString := "Hello Go"

	// greeting <- greetString // blocking because it is continuosly trying to
	// // receive value, it is ready to receive continuos flow of data

	go func() {
		greeting <- greetString
		greeting <- "greetString"

		for _, e := range "abcde" {
			greeting <- "Alphabet: " + string(e)
		}
	}()
	go func() {
		for range 7 {
			receiver := <-greeting
			fmt.Println(receiver)
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("End of Program")
}
