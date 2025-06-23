package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, GO standard library!")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error Fetching the data")
		return
	}
	defer resp.Body.Close()
	fmt.Println("HTTP Response Status: ", resp.Status)
}

/*
Using named Imports


package main

import (
	"fmt"
	foo "net/http"
)

func main(){
	fmt.Println("Hello, GO standard library!")

	resp, err  := foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error Fetching the data")
		return
	}
	defer resp.Body.Close()
	fmt.Println("HTTP Response Status: ", resp.Status)
}



*/
