package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age,omitempty"`
	Email   string   `xml:"-"`
	Address Address `xml:"address"`
}

type Address struct{
	City string `xml:"city"`
	State string `xml:"state"`
}



func main() {

	person := Person{Name: "John", Age: 30, Email: "email@example.com", Address: Address{City:"New York", State:"NY"}}

	xmlData, err := xml.MarshalIndent(person,"","   ")
	if err != nil {
		log.Fatalln("Error Marshalling data into XML:", err)
	}
	fmt.Println(string(xmlData))



	person2 := Person{Name: "John", Address: Address{City:"New York", State:"NY"}}


	xmlData, err = xml.MarshalIndent(person2,"","   ")
	if err != nil {
		log.Fatalln("Error Marshalling data into XML:", err)
	}
	fmt.Println(string(xmlData))


	// Handling raw data
	xmlRaw := `<person><name>James</name><age>52</age></person>`
	var personxml Person

	err = xml.Unmarshal([]byte(xmlRaw), &personxml)
	if err != nil {
		log.Fatalln("Error Unmarshalling XML into data:", err)
	}
	fmt.Println(personxml)
	fmt.Println("Local String:", personxml.XMLName.Local)
	fmt.Println("Namespace String:", personxml.XMLName.Space)

	book := Book{
		ISBN: "4538-7265-2473-859-604-645",
		Title: "Go Bootcamp",
		Author: "Jame Doe",
		Pseudo: "This is a pseudo message",
	}
	xmlDataAttr, err := xml.MarshalIndent(book, "", "   ")
	if err!= nil {
		log.Fatalln("Error mashalling book XML:", err)
	}
	fmt.Println(string(xmlDataAttr))
	

}


type Book struct {
	XMLName xml.Name `xml:"book"`
	ISBN string `xml:"isbn,attr"`
	Title string `xml:"title,attr"`
	Author string `xml:"author,attr"`
	Pseudo string `xml:"pseudo"`
}
