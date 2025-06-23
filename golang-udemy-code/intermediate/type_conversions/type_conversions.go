package main

import "fmt"

func main() {

	var a int = 32
	b := int32(a)
	c := float64(b)

	e := 3.14
	f := int(e)
	fmt.Println(f,c)

	// Type(value)

	g:= "Hello $#!;"
	
	var h []byte = []byte(g)
	fmt.Println(h)

	i := []byte{78,34,255}
	j := string(i)
	fmt.Println(j)
}