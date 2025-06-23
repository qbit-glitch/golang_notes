package main

import "fmt"

type Rectangle struct {
	length float64
	width float64
}

type Shape struct {
	Rectangle
}


// Method with value receiver
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Method with Pointer Receiver
func (r *Rectangle) Scale(factor float64){
	r.length *= factor
	r.width *= factor
}




func main() {

	rect := Rectangle{
		length:10, width:8,
	}
	area := rect.Area()
	fmt.Println("Area of rectangle with width 8 and length 10 is ", area)

	rect.Scale(2)
	area = rect.Area()
	fmt.Println("Area of rectangle after scaling by factor of 2 is", area)

	num := MyInt(7)
	num1 := MyInt(-7)
	fmt.Println(num.IsPositive())
	fmt.Println(num1.IsPositive())
	fmt.Println(num.welcomeMessage())

	s := Shape{Rectangle: Rectangle{length: 10, width:4}}
	fmt.Println(s.Area())
	fmt.Println(s.Rectangle.Area())

	/* Now we are embedding a struct inside another struct,
	the method associated with the embedded struct will be 
	promoted to the outer struct directly, so we can access
	 area directly.
	*/

}


type MyInt int

// Method on a user-defined type
func (m MyInt) IsPositive() bool {
	return m > 0
}

func (MyInt) welcomeMessage() string {
	return "Welcome to MyInt Type"
}

