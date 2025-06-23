package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("User env var:", user)
	fmt.Println("Home env var:", home)

	err := os.Setenv("FRUIT", "apple")
	if err!= nil {
		fmt.Println("error setting FRUIT env var:", err)
	} else {
		fmt.Println("Setting FRUIT env successful")
	}
	fmt.Println("FRUIT env var:", os.Getenv("FRUIT"))
	


	for _, e := range os.Environ(){
		kvpair := strings.SplitN(e,"=", 2)
		fmt.Println(kvpair[0])
	}

	err = os.Unsetenv("FRUIT")
	if err != nil {
		fmt.Println("error unsetting the environment var:", err)
		return
	}
	fmt.Println("Unset env var done on key FRUIT")
	fmt.Println("FRUIT env var:", os.Getenv("FRUIT"))

	/*
	"a=b=c=d"
	n=1 -> returns "a=b=c=d"
	n=2 -> returns "a" and "b=c=d"
	n=3 -> returns "a" and "b" and "c=d"
	n=4 -> returns "a" and "b" and "c" and "d"
	*/
	str := "a=b=c=d"
	fmt.Println(strings.SplitN(str, "=", -1))
	fmt.Println(strings.SplitN(str, "=", 0))
	fmt.Println(strings.SplitN(str, "=", 1))
	fmt.Println(strings.SplitN(str, "=", 2))
	fmt.Println(strings.SplitN(str, "=", 3))
	fmt.Println(strings.SplitN(str, "=", 4))

}