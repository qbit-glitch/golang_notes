package main

import "fmt"

func init(){
	fmt.Println("Initializing Package1 ...")
}
func init(){
	fmt.Println("Initializing Package2 ...")
}
func init(){
	fmt.Println("Initializing Package3 ...")
}

func init(){
	fmt.Println("Initializing Package4 ...")
}

func main() {

	fmt.Println("inside the main function")

}