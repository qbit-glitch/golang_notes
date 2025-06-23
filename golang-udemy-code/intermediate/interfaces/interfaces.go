package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type rect1 struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}


func (r rect1) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

func (r rect1) perim() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) dia() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area(), g.perim())
}

func main() {

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	r1 := rect1{width: 3, height: 4}

	measure(r)
	measure(c)
	measure(r1)

	myPrinter(1, "John", 45.9, true)

	printType(1)
	printType("John")
	printType(false)
}

func printType(i interface{}){    // can accept any type of value but only 1 number of parameters
	switch i.(type) {
	case int:
		fmt.Println("Type: Int")
	case string:
		fmt.Println("Type: string")
	default:
		fmt.Println("Type: Unknown")
	}
}



func myPrinter(i ...interface{}){   // can accept any number of parameters and any types of values
	for _, v := range i {
		fmt.Println(v)
	}
}