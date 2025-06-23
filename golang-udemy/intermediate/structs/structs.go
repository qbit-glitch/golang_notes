package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	PhoneHomeCell      // Embedding an Anonymous Field
}

type Address struct {
	city    string
	country string
}

type PhoneHomeCell struct {
	home string
	cell string
}

func main() {

	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		address: Address{
			city:    "London",
			country: "U.K",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "oveaoveirb",
			cell: "3456789876543",
		},
	}

	p1 := Person{
		firstName: "Smith",
		age:       30,

	}

	p3 := Person{
		firstName: "Smith",
		age:       30,

	}


	// p1.address.city = "New York"
	// p1.address.country = "USA"
	// p1.cell = "8765434567"
	// p1.home = "TEvuvcwe"

	fmt.Println(p)
	fmt.Println(p1)

	fmt.Println(p.firstName)
	fmt.Println(p1.firstName)

	fmt.Println(p.fullName())
	fmt.Println(p.cell)

	fmt.Println("Are p and p1 equal:", p == p1)
	fmt.Println("Are p3 and p1 equal:", p3 == p1)

	// Anonymous Struct
	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "pseudoemail@example.com",
	}

	fmt.Println(user)

	fmt.Println("Before Increment:", p1.age)
	p1.incrementAgeByOne()
	fmt.Println("After Increment:", p1.age)
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}
